import * as k8s from "@pulumi/kubernetes";
// import * as tls from "@pulumi/tls";
import * as config from "../../config";
import * as targets from "../../targets";
import * as primitives from "../00-primitives/mod";

const allowedRoutes = {
  namespaces: {
    from: "Selector",
    selector: {
      matchLabels: {
        "shared-gateway-access": "true",
      },
    },
  },
};

// custom resources

export const gatewayCrds = new k8s.yaml.ConfigFile(
  "gateway-api",
  {
    file:
      "https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.0.0/standard-install.yaml",
  },
  { provider: targets.k8sProvider },
);

// certificates

// export const sharedTlsCerts = new tls.PrivateKey("shared-tls-certs", {
//   algorithm: "RSA",
//   rsaBits: 4096,
// });

// const privateKey = sharedTlsCerts.privateKeyPemPkcs8.apply((k) =>
//   Buffer.from(k).toString("base64")
// );

// const publicKey = sharedTlsCerts.publicKeyPem.apply((k) =>
//   Buffer.from(k).toString("base64")
// );

const privateKey = config.privateKey;
const publicKey = config.publicKey;

// secrets

export let sharedTls: k8s.core.v1.Secret | undefined = undefined;

if (privateKey !== undefined && publicKey !== undefined) {
  sharedTls = new k8s.core.v1.Secret("shared-tls", {
    metadata: {
      name: "shared-tls",
    },
    type: "tls",
    data: {
      "tls.key": privateKey,
      "tls.crt": publicKey,
    },
  });
}

// gateways

const certificateRefs = [];

if (sharedTls !== undefined) {
  certificateRefs.push({ name: sharedTls.metadata.name });
}

export const gateway = new k8s.apiextensions.CustomResource(
  "gateway",
  {
    apiVersion: "gateway.networking.k8s.io/v1beta1",
    kind: "Gateway",
    metadata: {
      name: "shared-gateway",
      namespace: primitives.defaultNs.metadata.name,
    },
    spec: {
      gatewayClassName: "istio",
      listeners: [
        {
          name: "http",
          port: 80,
          protocol: "HTTP",
          allowedRoutes: allowedRoutes,
        },
        {
          name: "https",
          port: 443,
          protocol: "HTTPS",
          tls: {
            mode: "Terminate",
            certificateRefs: certificateRefs,
          },
          allowedRoutes: allowedRoutes,
        },
      ],
    },
  },
  { provider: targets.k8sProvider },
);
