# Edit a Predefined Database Service and apply changes only on redeploy

Status: easy-ai

## What to build

Add the edit path for **Predefined Database Services** so the user can change saved configuration without silently mutating the live runtime. Runtime changes should apply only after an explicit redeploy through the normal dashboard flow.

## Acceptance criteria

- [ ] A user can edit the supported saved fields for a **Predefined Database Service** from the dashboard.
- [ ] Saving changes updates stored configuration without immediately mutating the running database service.
- [ ] The user can trigger an explicit redeploy to apply the edited settings to runtime.

## Blocked by

- `.scratch/demo-stable-v1/issues/15-create-a-postgres-predefined-database-service-with-an-internal-url.md`
