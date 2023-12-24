import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";

// namespace

const nsName = "argowf";
export const ns = new k8s.core.v1.Namespace(nsName, {
  metadata: { name: nsName },
});

// helm chart

export const chart = new k8s.helm.v3.Chart("argowf", {
  namespace: ns.metadata.name,
  chart: "argo-workflows",
  fetchOpts: { repo: "https://argoproj.github.io/argo-helm" },
  values: {
    installCRDs: true,
    server: {
      serviceType: config.useLoadBalancer ? "LoadBalancer" : "ClusterIP",
    },
  },
});
