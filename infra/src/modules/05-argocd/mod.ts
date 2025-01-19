// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";
import * as targets from "../../targets";
import * as gateway from "../03-gateway/mod";

// namespaces

const nsName = "argocd";
export const ns = new k8s.core.v1.Namespace(
  "argocd-namespace",
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

export const chart = new k8s.helm.v3.Release(
  "argocd-helm-chart",
  {
    name: "argocd",
    namespace: ns.metadata.name,
    chart: "argo-cd",
    repositoryOpts: { repo: "https://argoproj.github.io/argo-helm" },
    values: {
      installCRDs: true,
      configs: {
        params: {
          "server.insecure": true,
        },
      },
    },
  },
  { provider: targets.k8sProvider, dependsOn: [ns] },
);

// http routes

export const httpRoute = new k8s.apiextensions.CustomResource(
  "argocd-http-route",
  {
    apiVersion: "gateway.networking.k8s.io/v1beta1",
    kind: "HTTPRoute",
    metadata: {
      name: "argocd",
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
        `cd.${config.domain}`,
      ],
      rules: [
        {
          backendRefs: [
            {
              name: "argocd-server",
              namespace: ns.metadata.name,
              port: 80,
            },
          ],
        },
      ],
    },
  },
  { provider: targets.k8sProvider, dependsOn: [ns] },
);
