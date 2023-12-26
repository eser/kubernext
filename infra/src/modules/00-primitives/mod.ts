import * as k8s from "@pulumi/kubernetes";
import * as targets from "../../targets";

// namespaces

const defaultNsName = "default";
export const defaultNs = k8s.core.v1.Namespace.get(
  "namespace",
  defaultNsName,
  { provider: targets.k8sProvider },
);
