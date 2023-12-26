import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";

// namespaces

const nsNameSystem = "istio-system";
export const nsSystem = new k8s.core.v1.Namespace(
  nsNameSystem,
  {
    metadata: {
      name: nsNameSystem,
    },
  },
  { provider: targets.k8sProvider },
);

const nsNameIngress = "istio-ingress";
export const nsIngress = new k8s.core.v1.Namespace(
  nsNameIngress,
  {
    metadata: {
      name: nsNameIngress,
      labels: {
        "istio-injection": "enabled",
      },
    },
  },
  { provider: targets.k8sProvider },
);

// helm charts

export const chartBase = new k8s.helm.v3.Chart(
  "istio-base",
  {
    namespace: nsSystem.metadata.name,
    chart: "base",
    fetchOpts: { repo: "https://istio-release.storage.googleapis.com/charts" },
  },
  { provider: targets.k8sProvider },
);

export const chartD = new k8s.helm.v3.Chart(
  "istiod",
  {
    namespace: nsSystem.metadata.name,
    chart: "istiod",
    fetchOpts: { repo: "https://istio-release.storage.googleapis.com/charts" },
  },
  { provider: targets.k8sProvider, dependsOn: [chartBase] },
);

export const chartGateway = new k8s.helm.v3.Chart(
  "istio-gateway",
  {
    namespace: nsIngress.metadata.name,
    chart: "gateway",
    fetchOpts: { repo: "https://istio-release.storage.googleapis.com/charts" },
  },
  { provider: targets.k8sProvider, dependsOn: [chartD] },
);
