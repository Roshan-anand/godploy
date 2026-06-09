# extras-v1 — 03: Project Instance Rename

## What to build

Allow renaming any **Project Instance** — both production and preview instances.

In the instance view (header or settings area), show an editable name field or a rename action. On save, the backend updates the instance name. Validate uniqueness within the same project (two instances cannot share a name).

Backend: PATCH endpoint on the instance resource to update the name field.

## Acceptance criteria

- [ ] Production instance name is editable
- [ ] Preview instance name is editable
- [ ] Name uniqueness enforced within a project
- [ ] Changes persist and display correctly in the instance switcher and breadcrumbs

## Blocked by

None — can start immediately.
