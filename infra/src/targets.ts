import * as k8s from "@pulumi/kubernetes";
import * as fs from "fs";
import * as config from "./config";

// k8s provider
let k8sProviderInstance: k8s.Provider;

if (config.kubeconfigPath) {
  const kubeconfig = fs.readFileSync(config.kubeconfigPath).toString();

  k8sProviderInstance = new k8s.Provider("provider", {
    kubeconfig: kubeconfig,
  });
} else {
  k8sProviderInstance = new k8s.Provider("provider", {});
}

export const k8sProvider = k8sProviderInstance;
