// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as k8s from "@pulumi/kubernetes";
import * as config from "../../config";
import * as targets from "../../targets";
import * as primitives from "../00-primitives/mod";

// metallb manifest

let metallbManifest: k8s.yaml.v2.ConfigFile | null = null;
let ipAddressPool: k8s.apiextensions.CustomResource | null = null;
let l2Advertisement: k8s.apiextensions.CustomResource | null = null;

if (config.installLoadBalancer) {
	metallbManifest = new k8s.yaml.v2.ConfigFile(
		"metallb-manifest",
		{
			file: `${config.cwd}/src/modules/01-load-balancer/metallb-manifest.yaml`,
		},
		{ provider: targets.k8sProvider, dependsOn: [primitives.defaultNs] },
	);

  if (config.loadBalancerAddressPool !== undefined) {
    ipAddressPool = new k8s.apiextensions.CustomResource(
      "default-address-pool",
      {
        apiVersion: "metallb.io/v1beta1",
        kind: "IPAddressPool",
        metadata: {
          name: "default-address-pool",
          namespace: "metallb-system",
        },
        spec: {
          addresses: [config.loadBalancerAddressPool],
        },
      },
      { provider: targets.k8sProvider, dependsOn: [metallbManifest] },
    );

    l2Advertisement = new k8s.apiextensions.CustomResource(
      "default-l2-advertisement",
      {
        apiVersion: "metallb.io/v1beta1",
        kind: "L2Advertisement",
        metadata: {
          name: "default",
          namespace: "metallb-system",
        },
        spec: {},
      },
      { provider: targets.k8sProvider, dependsOn: [metallbManifest] },
    );
  }
}

export { metallbManifest, ipAddressPool, l2Advertisement };
