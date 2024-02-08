# KuberNeXT Infra

KuberNeXT infrastructure as code project / uses Pulumi

## Deployment

To deploy your infrastructure, follow the steps below.

### Prerequisites

1. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)
2. Install NPM Dependencies
3. [Install `kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### Steps

After cloning this repo, run these commands from the working directory:

1. Login to pulumi:

   ```bash
   $ pulumi login [your-pulumi-state-bucket-url]
   ```

2. Execute the Pulumi program to create or update your infra:

   ```bash
   $ pulumi up
   ```

### Argo CD

1. Visit `http://localhost`
2. Get password with
   `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`
   command
3. Login as admin with the password

### Argo Workflow

1. Argo port forward

   ```bash
   $ kubectl -n argowf port-forward deployment/argo-server 2746:2746
   ```
