import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";
import * as gateway from "../02-gateway/mod";

// namespaces

const nsName = "argowf";
export const ns = new k8s.core.v1.Namespace(
  nsName,
  {
    metadata: {
      name: nsName,
      labels: {
        "shared-gateway-access": "true",
      },
    },
  },
  { provider: targets.k8sProvider },
);

// helm charts

export const chart = new k8s.helm.v3.Chart(
  "argowf",
  {
    namespace: ns.metadata.name,
    chart: "argo-workflows",
    fetchOpts: { repo: "https://argoproj.github.io/argo-helm" },
    values: {
      installCRDs: true,
    },
  },
  { provider: targets.k8sProvider },
);

// http routes

export const httpRoute = new k8s.apiextensions.CustomResource(
  "argowf",
  {
    apiVersion: "gateway.networking.k8s.io/v1beta1",
    kind: "HTTPRoute",
    metadata: {
      name: "argowf",
      namespace: ns.metadata.name,
    },
    spec: {
      parentRefs: [
        {
          name: gateway.gateway.metadata.name,
          namespace: gateway.gateway.metadata.namespace,
        },
      ],
      hostnames: [
        "workflows.eser.land",
      ],
      rules: [
        {
          backendRefs: [
            {
              name: "argowf-argo-workflows-server",
              namespace: ns.metadata.name,
              port: 2746,
            },
          ],
        },
      ],
    },
  },
  { provider: targets.k8sProvider },
);
