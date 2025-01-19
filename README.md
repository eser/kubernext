# Project KuberNeXT

This project promises to standardize the way of working. While offering modern
structure for every software team, also enforces a well-defined structure by
setting some rules.

## Documentation

- [Proposed Workflow for Changes](./docs/workflow.md)
- [Proposed Tools](./docs/tools.md)

## Infrastructure

This project uses [Infrastructure as Code](./infra/) with Pulumi. Current stack consists of:

- [kind](https://kind.sigs.k8s.io/) for local Kubernetes cluster
- [Pulumi](https://www.pulumi.com/) for infrastructure provisioning
- [Argo Workflows](https://argoproj.github.io/workflows) for worker tasks
- [ArgoCD](https://argoproj.github.io/cd) for GitOps
- [Prometheus](https://prometheus.io/) for metrics
- [Grafana](https://grafana.com/) for monitoring
- [Istio](https://istio.io/) for service mesh
- [MetalLB](https://metallb.universe.tf/) for load balancer

## Decision Logs

"Architectural Design Records" format is being used under this repository. See
[docs/adr/](./docs/adr/) for the decisions and proposals recorded so far.

## Codebase

| Directory    | Description                  | Link               |
| ------------ | ---------------------------- | ------------------ |
| `infra/`     | Infrastructure as Code       | [Go](./infra/)     |
| `docs/`      | Documentation                | [Go](./docs/)      |
| `docs/adr/`  | Architectural Design Records | [Go](./docs/adr/)  |
| `templates/` | Project boilerplates         | [Go](./templates/) |
| `apps/`      | App packages ready to deploy | [Go](./apps/)      |
