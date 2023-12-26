import "./modules/01-istio/mod";
import "./modules/02-gateway/mod";
import "./modules/03-argowf/mod";
import "./modules/04-argocd/mod";
import "./modules/05-test-nginx/mod";

import * as config from "./config";

console.log(`using kubeconfig: ${config.kubeconfigPath}`);
