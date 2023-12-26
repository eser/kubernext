import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";

// namespaces

const nsName = "argocd";
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
  "argocd",
  {
    namespace: ns.metadata.name,
    chart: "argo-cd",
    fetchOpts: { repo: "https://argoproj.github.io/argo-helm" },
    values: {
      installCRDs: true,
    },
  },
  { provider: targets.k8sProvider },
);
