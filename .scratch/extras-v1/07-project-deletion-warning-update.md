# extras-v1 — 07: Project Deletion — Warning Update

## What to build

Update the existing project deletion confirmation dialog. Currently it warns about orphan volumes. Change it to warn about running services that will be removed.

When the user clicks "Delete Project":
- Show a confirmation dialog listing all services (application + predefined DB) that are currently running or have active deployments across all instances.
- The dialog should name each service and which instance it belongs to.
- The dialog should also note that all data will be lost (orphan volume handling is separate — the user should transfer/preserve volumes before deletion).
- Remove the orphan volume warning from this flow (it has its own handling in the Storage area).

This is a UI-only change. Backend behavior remains the same.

## Acceptance criteria

- [ ] Delete project dialog shows a list of running services per instance
- [ ] Dialog clearly states that all data will be removed
- [ ] No orphan volume warnings in the project delete flow
- [ ] Confirmation still requires explicit user action before deletion proceeds

## Blocked by

None — can start immediately.
