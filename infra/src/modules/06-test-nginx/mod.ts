// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";
import * as argocd from "../05-argocd/mod";

const appName = "nginx-test";

// argocd application

export const application = new k8s.apiextensions.CustomResource(
  appName,
  {
    apiVersion: "argoproj.io/v1alpha1",
    kind: "Application",
    metadata: {
      name: appName,
      namespace: argocd.ns.metadata.name,
    },
    spec: {
      project: "default",
      source: {
        repoURL: "https://github.com/eser/kubernext.git",
        targetRevision: "HEAD",
        path: "apps/nginx-test/deployments/k8s/overlays/default",
      },
      destination: {
        server: "https://kubernetes.default.svc",
        namespace: "nginx-test",
      }
    },
  },
  { provider: targets.k8sProvider },
);
