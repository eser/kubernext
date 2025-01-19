// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import * as pulumi from "@pulumi/pulumi";

export const cwd = process.cwd();

export const instance = new pulumi.Config();

export const installLoadBalancer = instance.getBoolean("installLoadBalancer");
export const loadBalancerAddressPool = instance.get("loadBalancerAddressPool");

export const domain = instance.require("domain");

// grafana

export const grafanaUsername = instance.require("grafanaUsername");
export const grafanaPassword = instance.require("grafanaPassword");

// keys

export const privateKey = instance.get("privateKey");
export const publicKey = instance.get("publicKey");
