import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";
import * as primitives from "../00-primitives/mod";
import * as gateway from "../02-gateway/mod";

const appName = "nginx";
const appLabels = { target: `${appName}-deployment` };

// deployments

export const deployment = new k8s.apps.v1.Deployment(
  appName,
  {
    metadata: {
      name: appName,
      namespace: primitives.defaultNs.metadata.name,
    },
    spec: {
      selector: { matchLabels: appLabels },
      replicas: 1,
      template: {
        metadata: { labels: appLabels },
        spec: { containers: [{ name: appName, image: "nginx" }] },
      },
    },
  },
  { provider: targets.k8sProvider },
);

// services

export const service = new k8s.core.v1.Service(
  appName,
  {
    metadata: {
      name: appName,
      namespace: primitives.defaultNs.metadata.name,
      labels: deployment.spec.template.metadata.labels,
    },
    spec: {
      type: "ClusterIP",
      ports: [{ port: 80, targetPort: 80, protocol: "TCP" }],
      selector: appLabels,
    },
  },
  { provider: targets.k8sProvider },
);

// http routes

export const httpRoute = new k8s.apiextensions.CustomResource(
  appName,
  {
    apiVersion: "gateway.networking.k8s.io/v1beta1",
    kind: "HTTPRoute",
    metadata: {
      name: appName,
      namespace: primitives.defaultNs.metadata.name,
    },
    spec: {
      parentRefs: [
        {
          name: gateway.gateway.metadata.name,
          namespace: gateway.gateway.metadata.namespace,
        },
      ],
      rules: [
        {
          matches: [
            {
              path: {
                type: "PathPrefix",
                value: "/test/",
              },
            },
          ],
          filters: [
            {
              type: "URLRewrite",
              urlRewrite: {
                path: {
                  type: "ReplacePrefixMatch",
                  replacePrefixMatch: "/",
                },
              },
            },
          ],
          backendRefs: [
            {
              name: service.metadata.name,
              namespace: service.metadata.namespace,
              port: 80,
            },
          ],
        },
      ],
    },
  },
  { provider: targets.k8sProvider },
);
