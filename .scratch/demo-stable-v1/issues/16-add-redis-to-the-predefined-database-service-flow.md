# Add Redis to the Predefined Database Service flow

Status: easy-ai

## What to build

Extend the predefined-service path so Redis is available as a second built-in **Predefined Database Service** template. The user should be able to create Redis through the same end-to-end flow style as Postgres and receive the right internal-only runtime behavior for the selected instance.

## Acceptance criteria

- [ ] Redis is available as a built-in **Predefined Database Service** template in the create flow.
- [ ] A user can deploy Redis from the normal flow and have it run as an internal-only service inside the selected instance.
- [ ] The dashboard treats Redis as part of the same predefined-service experience rather than as a one-off special case.

## Blocked by

- `.scratch/demo-stable-v1/issues/15-create-a-postgres-predefined-database-service-with-an-internal-url.md`
