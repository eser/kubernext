# Proposed Tools

- [Pulumi](https://www.pulumi.com/) for infrastructure provisioning
- [Kubernetes](https://kubernetes.io/) for container orchestration
- [Argo Workflows](https://argoproj.github.io/workflows) for worker tasks
- [ArgoCD](https://argoproj.github.io/cd) for GitOps

## Possible Ambiguity

If you want to incorporate an application into your infrastructure and it's an
application you did not package yourself (for example, if it is installed with
Helm), then in 90% of the cases, you should create it using the Infrastructure
as Code approach with Pulumi.

In-house applications that you have packaged yourself can be installed via the
interface by pointing to the target repository through ArgoCD, using the GitOps
approach.
