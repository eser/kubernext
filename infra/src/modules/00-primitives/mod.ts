// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";

// namespaces

const defaultNsName = "default";
export const defaultNs = k8s.core.v1.Namespace.get(
	"default-namespace",
	defaultNsName,
	{ provider: targets.k8sProvider },
);
