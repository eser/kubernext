// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";

// k8s provider
export const k8sProvider = new k8s.Provider("provider", {});
