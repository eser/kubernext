import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";

// namespace

const nsName = "argocd";
export const ns = new k8s.core.v1.Namespace(nsName, {
  metadata: { name: nsName },
});

// helm chart

export const chart = new k8s.helm.v3.Chart("argocd", {
  namespace: ns.metadata.name,
  chart: "argo-cd",
  fetchOpts: { repo: "https://argoproj.github.io/argo-helm" },
  values: {
    installCRDs: true,
    server: {
      service: {
        type: config.useLoadBalancer ? "LoadBalancer" : "ClusterIP",
      },
    },
  },
});
