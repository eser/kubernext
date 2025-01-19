# Project KuberNeXT

KuberNeXT is a powerful tool designed to deploy a meticulously crafted DevOps stack in minutes, whether on a remote server or your local development environment. This stack includes fully-featured tools for GitOps, monitoring, metrics, and workflows—just like the robust infrastructure you'd find in a leading software company.

## Infrastructure

This project uses [Infrastructure as Code](./infra/) with Pulumi. Current stack consists of:

- [Kubernetes](https://kubernetes.io/) for container orchestration
- [kind](https://kind.sigs.k8s.io/) for creating local Kubernetes cluster(s)
- [Pulumi](https://www.pulumi.com/) for infrastructure provisioning
- [Argo Workflows](https://argoproj.github.io/workflows) for worker tasks
- [ArgoCD](https://argoproj.github.io/cd) for GitOps
- [Prometheus](https://prometheus.io/) for metrics
- [Grafana](https://grafana.com/) for monitoring
- [Istio](https://istio.io/) for service mesh
- [MetalLB](https://metallb.universe.tf/) for load balancer

## Deployment

To deploy your infrastructure, follow the steps below.

### Prerequisites

1. [Install kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) for local Kubernetes cluster
2. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)
3. Install NPM Dependencies
4. [Install `kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### Steps

After cloning this repo, run these commands from the working directory:


#### kind

```bash
$ kind create cluster --name local --config kind-config.yaml
```

#### Pulumi

1. Login to pulumi:

   ```bash
   $ pulumi login --local
   ```

2. Create a new stack:

   ```bash
   $ pulumi stack init production
   ```

2. Execute the Pulumi program to create or update your infra:

   ```bash
   $ pulumi up --stack production
   ```

#### Argo CD

1. Visit `http://localhost`
2. Get password with
   `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`
   command
3. Login as admin with the password

#### Argo Workflow

1. Argo port forward

   ```bash
   $ kubectl -n argowf port-forward deployment/argo-server 2746:2746
   ```

### License

This project is licensed under the Apache 2.0 License. For further details,
please see the [LICENSE](LICENSE) file.

### To support the project...

[Visit my GitHub Sponsors profile at github.com/sponsors/eser](https://github.com/sponsors/eser)
