# Reattach a compatible Orphan Volume during Predefined Database Service creation

Status: need-testing

## What to build

Extend predefined database creation so the user can attach a compatible **Orphan Volume** instead of always starting from a fresh volume. The flow should keep restore behavior visible and bounded to compatible predefined-service types.

## Acceptance criteria

- [x] A user can select a compatible **Orphan Volume** while creating a matching **Predefined Database Service**.
- [x] Reattaching the volume removes it from **Storage** and assigns it to the new service.
- [x] Incompatible or risky restores are handled through the normal UI flow with the intended warnings or constraints.

## Blocked by

- `.scratch/demo-stable-v1/issues/18-preserve-deleted-database-data-as-an-orphan-volume-and-show-it-in-storage.md`
