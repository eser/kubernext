# KuberNeXT Infra

KuberNeXT infrastructure as code project / uses Pulumi

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
