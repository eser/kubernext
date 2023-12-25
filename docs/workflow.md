# Proposed Workflow for Changes

## Planning

Plan using your own method. The output of this planning should be PBIs
(features, improvements, bug fixes, etc.). Then, create a release schedule with
semantic versions.

To measure, ensure that each PBI includes:

- Target version
- Acceptance criteria
- Quality criteria
- Test cases and scenarios
- Business impact value

## Implementation / Coding

Code changes are made on a branch associated with a PBI (something like
feature/dynamic-pricing). Frequent commits should be made and a Pull Request
should be opened to the public branch for code review and tests.

## Testing

TBD

## Merging

The development PR, which is verified to be working and has passed the checks,
should be squash merged into the public branch. Ensure that the commit messages
created by the squash merge comply with the conventional commit messages
standard. The changes made should be recorded in the Release Log, along with the
PBI ID and commit ID.

## Releasing / Delivery

A release is created on SCM and a Tag is generated at the relevant point in the
code. This tag should follow the "v0.0.0" format and adhere to the rules of
semantic versioning. The Release Log generated during the squash merge should
also be added to the description of the release.

The creation of a release should automatically trigger the generation of a
containerized build of the project with its current version tag in the
artifact/image registry.

The build found in the artifact/image registry should be ready to be deployed in
any environment.

## Deployment

The release of a new version will also be reflected in our CD tool (i.e.
ArgoCD), which facilitates our work with the GitOps approach. Once the new
version is approved, the implementation of changes can also be managed through
the tool's own interface.
