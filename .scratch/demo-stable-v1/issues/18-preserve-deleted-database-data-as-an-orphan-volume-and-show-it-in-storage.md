# Preserve deleted database data as an Orphan Volume and show it in Storage

Status: need-testing

## What to build

Complete the delete flow for **Predefined Database Services** so preserved data becomes an **Orphan Volume** visible in **Storage**. The user should be able to remove the runtime while keeping the data as an explicit follow-up resource.

## Acceptance criteria

- [x] Deleting a predefined database service can preserve its data instead of purging it.
- [x] Preserved data becomes an **Orphan Volume** rather than staying attached to the deleted service.
- [x] The **Storage** area shows preserved orphan volumes through the normal dashboard flow.

## Blocked by

- `.scratch/demo-stable-v1/issues/15-create-a-postgres-predefined-database-service-with-an-internal-url.md`
