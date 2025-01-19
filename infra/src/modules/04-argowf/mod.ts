// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";
import * as targets from "../../targets";
import * as gateway from "../03-gateway/mod";

// namespaces

const nsName = "argowf";
export const ns = new k8s.core.v1.Namespace(
  "argowf-namespace",
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
  "argowf-helm-chart",
  {
    name: "argowf",
    namespace: ns.metadata.name,
    chart: "argo-workflows",
    repositoryOpts: { repo: "https://argoproj.github.io/argo-helm" },
    values: {
      installCRDs: true,
    },
  },
  { provider: targets.k8sProvider, dependsOn: [ns] },
);

// http routes

export const httpRoute = new k8s.apiextensions.CustomResource(
  "argowf-http-route",
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
        `workflows.${config.domain}`,
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
  { provider: targets.k8sProvider, dependsOn: [ns] },
);
