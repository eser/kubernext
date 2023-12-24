# 4. Install Kubernetes

Date: 2023-10-25\
Champion: Eser Ozvataf

## Status

Accepted

## Context

We need to remove boundaries and frictions of our developers as they are
executing their daily tasks to speed up our development process.

Therefore, our developments environments should be portable and can be run on
server and/or other places without any difference. Additionally, there shouldn't
be any vendor lock-in in terms of being platform-agnostic.

## Decision

We will use _containerization_ and Kubernetes to orchestrate these containers.

## Consequences

- **Kubernetes** will be installed, (see: https://kubernetes.io/).
- All solutions will be bundled/dockerized.
