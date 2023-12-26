import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";
import * as primitives from "../00-primitives/mod";

// custom resources

export const gatewayCrds = new k8s.yaml.ConfigFile(
  "gateway-api",
  {
    file:
      "https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.0.0/standard-install.yaml",
  },
  { provider: targets.k8sProvider },
);

// gateways

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
        },
        {
          name: "https",
          port: 443,
          protocol: "HTTPS",
          tls: {
            mode: "Terminate",
            certificateRefs: [
              {
                kind: "Secret",
                group: "",
                name: "shared-tls",
              },
            ],
          },
        },
      ],
    },
  },
  { provider: targets.k8sProvider },
);
