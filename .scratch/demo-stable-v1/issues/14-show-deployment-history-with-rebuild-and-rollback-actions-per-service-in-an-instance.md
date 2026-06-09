# Show deployment history with rebuild and rollback actions per Service in an instance

Status: easy-ai

## What to build

Deliver service-level deployment visibility and control inside both production and preview instances. Each application **Service** should show its deployment history and support rebuild and rollback actions through normal dashboard flows that are safer around edge cases.

## Acceptance criteria

- [ ] A user can view deployment history for an individual application **Service** inside the selected instance.
- [ ] A user can trigger rebuild and rollback actions for that service through the normal UI flow.
- [ ] The rebuild and rollback paths handle the expected edge cases cleanly enough to avoid demo-breaking surprises.

## Blocked by

- `.scratch/demo-stable-v1/issues/03-create-an-internal-application-service-inside-the-production-project-instance.md`
