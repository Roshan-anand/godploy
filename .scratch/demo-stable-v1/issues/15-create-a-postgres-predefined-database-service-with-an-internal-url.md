# Create a Postgres Predefined Database Service with an Internal URL

Status: need-testing

## What to build

Deliver the full logged-in flow for creating a Postgres **Predefined Database Service** inside the selected **Project Instance**. The flow should let the user choose the instance context, start from the built-in Postgres template, edit the safe fields, deploy the database, and receive the generated **Internal URL** for manual application wiring.

## Acceptance criteria

- [x] The create-service flow lets the user create a Postgres **Predefined Database Service** from the built-in template with editable safe fields.
- [x] The resulting Postgres service is internal-only and deploys inside the selected instance network.
- [x] The user can view the generated full **Internal URL** from the normal dashboard flow.

## Blocked by

- `.scratch/demo-stable-v1/issues/02-create-a-project-with-a-default-production-project-instance.md`
