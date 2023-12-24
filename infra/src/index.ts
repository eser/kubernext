// import * as k8s from "@pulumi/kubernetes";
import * as argocd from "./modules/01-argocd/mod";
import * as argowf from "./modules/02-argowf/mod";
// import * as config from "./config";

// const appName = "nginx";
// const appLabels = { app: appName };
// const deployment = new k8s.apps.v1.Deployment(appName, {
//   spec: {
//     selector: { matchLabels: appLabels },
//     replicas: 1,
//     template: {
//       metadata: { labels: appLabels },
//       spec: { containers: [{ name: appName, image: "nginx" }] },
//     },
//   },
// });
// export const name = deployment.metadata.name;

// // Allocate an IP to the Deployment.
// const frontend = new k8s.core.v1.Service(appName, {
//   metadata: { labels: deployment.spec.template.metadata.labels },
//   spec: {
//     type: config.useLoadBalancer ? "LoadBalancer" : "ClusterIP",
//     ports: [{ port: config.frontendPort, targetPort: 80, protocol: "TCP" }],
//     selector: appLabels,
//   },
// });

// // When "done", this will print the public IP.
// export const ip = config.useLoadBalancer
//   ? frontend.status.loadBalancer.apply(
//     (lb) => lb.ingress[0].ip ?? lb.ingress[0].hostname,
//   )
//   : frontend.spec.clusterIP;

export const nsArgoCD = argocd.ns.metadata.name;
export const nsArgoWF = argowf.ns.metadata.name;
