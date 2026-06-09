# Show preview instance lifecycle status separately from deployment status

Status: easy-ai

## What to build

Add explicit **Project Instance** lifecycle visibility so the user can distinguish whether a preview is creating, ready, updating, deleting, or in error independently from per-service deployment and runtime status. The result should make preview orchestration state visible without hiding the underlying service-level statuses.

## Acceptance criteria

- [ ] Preview **Project Instances** expose their own lifecycle status separate from per-service deployment status.
- [ ] The dashboard shows preview lifecycle state clearly while still letting the user inspect service-level status underneath.
- [ ] Failed or incomplete preview orchestration is visible as instance state rather than only as scattered service-level symptoms.

## Blocked by

- `.scratch/demo-stable-v1/issues/06-create-a-branch-preview-project-instance-from-a-production-snapshot.md`
