# Expose a public application Service from the production Project Instance

Status: easy-ai

## What to build

Extend production application service creation so a public **Service** gets external ingress while keeping membership in the production **Instance Network**. The production service should own the stable base domain that later preview domains derive from.

## Acceptance criteria

- [x] A user can create a public application **Service** from the normal flow and have it exposed through Traefik.
- [x] The production runtime keeps a stable base domain suitable for later preview-domain generation.
- [x] The dashboard clearly distinguishes public and internal application service behavior during and after creation.

## Blocked by

- `.scratch/demo-stable-v1/issues/03-create-an-internal-application-service-inside-the-production-project-instance.md`
