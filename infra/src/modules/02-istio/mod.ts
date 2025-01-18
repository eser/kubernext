// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";
import * as targets from "../../targets";
import * as primitives from "../00-primitives/mod";

// namespace

export const istioSystemNs = new k8s.core.v1.Namespace(
	"istio-system-namespace",
	{
		metadata: {
			name: "istio-system",
		},
	},
	{ provider: targets.k8sProvider },
);

// istio manifest

export const istioManifest = new k8s.yaml.v2.ConfigFile(
	"istio-manifest",
	{
		file: `${config.cwd}/src/modules/02-istio/istio-manifest.yaml`,
	},
	{
		provider: targets.k8sProvider,
		dependsOn: [primitives.defaultNs, istioSystemNs],
	},
);
