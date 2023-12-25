# Project KuberNeXT

This project promises to standardize the way of working. While offering modern
structure for every software team, also enforces a well-defined structure by
setting some rules.

## Proposed Workflow

### Planning

TBD

### Implementation / Coding

Code changes are made on a branch associated with a development item on GitHub.
Frequent commits are made, and a Pull Request is opened to the public branch for
code review and tests.

### Testing

TBD

### Merging

The development PR, which is verified to be working and has passed the checks,
is squash merged into the public branch. It is ensured that the commit messages
created by squash merge complies with the conventional commit messages standard.
The changes made are recorded in the Release Log along with the development item
ID and commit ID.

### Releasing / Delivery

A release is created on GitHub, and a Tag is generated at the relevant point in
the code. This tag follows the "v0.0.0" format and adheres to the rules of
semantic versioning. The Release Log generated during the squash merge is also
added to the description of the release.

The creation of a release automatically triggers the generation of a
containerized build of the project with its current version tag in the
artifact/image registry.

The build found in the artifact/image registry is ready to be deployed in any
environment.

### Deployment

The release of a new version will also be reflected in ArgoCD, which facilitates
our work with the GitOps approach. Once the new version is approved, the
implementation of changes can also be managed through the ArgoCD interface.

## Decision Logs

"Architectural Design Records" format is being used under this repository. See
[docs/adr/](./docs/adr/) for the decisions and proposals recorded so far.

## Codebase

| Directory    | Description                  | Link               |
| ------------ | ---------------------------- | ------------------ |
| `docs/adr/`  | Architectural Design Records | [Go](./docs/adr/)  |
| `infra/`     | Infrastructure as Code       | [Go](./infra/)     |
| `templates/` | Project boilerplates         | [Go](./templates/) |
