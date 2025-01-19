// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";
import * as targets from "../../targets";
import * as gateway from "../03-gateway/mod";

// namespaces

const nsName = "monitoring";
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

export const kubePrometheusStackChart = new k8s.helm.v3.Release(
  "kube-prometheus-stack-helm-chart",
  {
    name: "prom",
    namespace: ns.metadata.name,
    chart: "kube-prometheus-stack",
    repositoryOpts: { repo: "https://prometheus-community.github.io/helm-charts" },
    values: {
      installCRDs: true,
    },
  },
  { provider: targets.k8sProvider, dependsOn: [ns] },
);

// grafana secrets

const updatedData = {
  "admin-user": Buffer.from(config.grafanaUsername).toString("base64"),
  "admin-password": Buffer.from(config.grafanaPassword).toString("base64"),
};

export const grafanaSecret = new k8s.core.v1.Secret(
  "grafana-secret",
  {
    metadata: {
      annotations: {
        "pulumi.com/patchForce": "true",
      },
      name: "prom-grafana",
      namespace: ns.metadata.name,
    },
    type: "Opaque",
    data: updatedData,
  },
  { provider: targets.k8sProvider, dependsOn: [ns, kubePrometheusStackChart] },
);

// http routes

export const httpRoute = new k8s.apiextensions.CustomResource(
  "grafana-http-route",
  {
    apiVersion: "gateway.networking.k8s.io/v1beta1",
    kind: "HTTPRoute",
    metadata: {
      name: "grafana",
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
        `grafana.${config.domain}`,
      ],
      rules: [
        {
          backendRefs: [
            {
              name: "prom-grafana",
              namespace: ns.metadata.name,
              port: 80,
            },
          ],
        },
      ],
    },
  },
  { provider: targets.k8sProvider, dependsOn: [ns, kubePrometheusStackChart] },
);
