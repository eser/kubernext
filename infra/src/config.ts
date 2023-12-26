import * as path from "path";
import * as os from "os";
import * as pulumi from "@pulumi/pulumi";

export const instance = new pulumi.Config();

// Minikube does not implement services of type `LoadBalancer`; require
// the user to specify if we're running on minikube, and if so, create
// only services of type ClusterIP.
export const useLoadBalancer = instance.requireBoolean("useLoadBalancer");

export const frontendPort = instance.requireNumber("frontendPort");

// determine the path to the kubeconfig file. If not specified, we'll assume
// it's the default one
const kubeconfigPathSet = instance.get("kubeconfigPath") ?? "~/.kube/config";

export const kubeconfigPath = kubeconfigPathSet.startsWith("~/")
  ? path.join(os.homedir(), kubeconfigPathSet.substring(2))
  : kubeconfigPathSet;
