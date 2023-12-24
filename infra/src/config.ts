import * as pulumi from "@pulumi/pulumi";

export const instance = new pulumi.Config();

// Minikube does not implement services of type `LoadBalancer`; require
// the user to specify if we're running on minikube, and if so, create
// only services of type ClusterIP.
export const useLoadBalancer = instance.requireBoolean("useLoadBalancer");

export const frontendPort = instance.requireNumber("frontendPort");
