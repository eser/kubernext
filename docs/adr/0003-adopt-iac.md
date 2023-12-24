# 3. Adopt IaC (Infrastructure as Code)

Date: 2023-10-25\
Champion: Eser Ozvataf

## Status

Accepted

## Context

We need to be able to reproduce our environments as much as possible.

## Decision

We will use Infrastructure as Code methodology, as
[described in Wikipedia](https://en.wikipedia.org/wiki/Infrastructure_as_code).

## Consequences

- **Pulumi** will be installed, (see: https://www.pulumi.com/).
- All infrastructure changes will be coded and kept in our SCM repositories.
