# 5. Setup CI/CD/CD pipelines separately

Date: 2023-10-25\
Champion: Eser Ozvataf

## Status

Accepted

## Context

We need mature build and deployment processes.

## Decision

We will setup separate pipelines for:

- [Continuous Integration](https://circleci.com/continuous-integration/)
- [Continuous Delivery](https://continuousdelivery.com/)
- [Continuous Deployment](https://scaledagileframework.com/continuous-deployment/)

## Consequences

- **GitHub Actions** will be used for Continuous Integration and Continuous
  Delivery.
  - All pull requests will be checked **before** merge operations.
  - Continuous integration pipelines will be executed **right after** _public
    branch_ commits including merges.
  - Configured branches will produce artifacts and push them to artifact
    registries (those artifacts will mostly be docker images, especially for
    code repositories).
- **ArgoCD** will be used for Continuous Deployment.
