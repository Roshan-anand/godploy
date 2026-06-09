# Manage generated preview domains and manual overrides per preview Service

Status: easy-ai

## What to build

Complete the public preview routing flow for application **Services** inside preview instances. Public preview services should receive generated domains using the preview identity and service name, and users should be able to override a preview service domain manually without affecting production or sibling previews.

## Acceptance criteria

- [ ] A public preview service receives a generated domain using the `<service>.<preview>.<base_domain>` pattern with runtime-safe slugs.
- [ ] A user can manually edit a preview service domain from the dashboard.
- [ ] Editing a preview service domain affects only that preview service and does not modify production or other previews.

## Blocked by

- `.scratch/demo-stable-v1/issues/04-expose-a-public-application-service-from-the-production-project-instance.md`
- `.scratch/demo-stable-v1/issues/06-create-a-branch-preview-project-instance-from-a-production-snapshot.md`
