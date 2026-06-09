# Demo-Stable V1 Execution Order

This file turns the published issue set into a practical execution sequence.

## Recommended strategy

- Start with the three root AFK slices that create the platform backbone:
  - `01` Install Godploy on Ubuntu from GHCR
  - `02` Create a Project with a default production Project Instance
  - `15` Create a Postgres Predefined Database Service with an Internal URL
- Run the human decision slices early enough that they do not surprise later implementation:
  - `22` Decide the V1 critical-flow test gate and coverage order
  - `23` Decide the V1 rate-limiting policy for core routes and service exposure
- Treat the GitHub callback contract as an early risk item once install behavior is stable:
  - `21` Align GitHub App manifest, webhook endpoint, and public server URL behavior

## Wave order

### Wave 0: root decisions and bootstrap

1. `01` Install Godploy on Ubuntu from GHCR
2. `22` Decide the V1 critical-flow test gate and coverage order
3. `23` Decide the V1 rate-limiting policy for core routes and service exposure

Why first:
- `01` is the base environment contract for demo-stable V1.
- `22` and `23` are still decision tickets, but resolving them early reduces rework in later slices.

### Wave 1: project and base runtime model

4. `02` Create a Project with a default production Project Instance
5. `03` Create an internal application Service inside the production Project Instance
6. `15` Create a Postgres Predefined Database Service with an Internal URL
7. `21` Align GitHub App manifest, webhook endpoint, and public server URL behavior

Why here:
- `02` establishes the new project-instance runtime boundary.
- `03` and `15` establish the two main product tracks: application service and predefined database service.
- `21` should be settled before preview and webhook automation goes deeper.

### Wave 2: production UX and instance-aware navigation

8. `04` Expose a public application Service from the production Project Instance
9. `05` Show project instance switching in the dashboard

Why here:
- Public routing and instance-aware navigation are the base UX prerequisites for preview work.

### Wave 3: preview creation path

10. `08` Keep open pull request candidates in SQLite and show them in the dashboard
11. `06` Create a branch preview Project Instance from a production snapshot
12. `07` Create a pull request preview Project Instance from available PR candidates

Why here:
- `06` proves the core preview snapshot and clone model.
- `07` builds on the same preview model but adds PR-specific source selection and candidate flow.
- `08` can proceed in parallel once the GitHub contract is settled.

### Wave 4: preview lifecycle and observability

13. `09` Update preview Project Instances from push and pull request events
14. `10` Auto-delete pull request preview Project Instances on close or merge
15. `11` Support manual delete and TTL cleanup for branch preview Project Instances
16. `12` Manage generated preview domains and manual overrides per preview Service
17. `13` Show preview instance lifecycle status separately from deployment status
18. `14` Show deployment history with rebuild and rollback actions per Service in an instance

Why here:
- These slices all deepen the preview lifecycle after preview creation exists.
- `12` depends on both public production routing and preview instance creation.
- `13` and `14` make the instance-level model understandable in the dashboard.

### Wave 5: predefined database follow-through

19. `16` Add Redis to the Predefined Database Service flow
20. `17` Edit a Predefined Database Service and apply changes only on redeploy
21. `18` Preserve deleted database data as an Orphan Volume and show it in Storage
22. `19` Reattach a compatible Orphan Volume during Predefined Database Service creation
23. `20` Warn about Orphan Volumes when deleting a Project

Why last:
- These all build directly on the first Postgres tracer bullet in `15`.
- `18` is the storage prerequisite for restore and project-delete warning flows.

## Parallel lanes

If multiple agents are available, use these lanes:

### Lane A: platform and GitHub contract

- `01` -> `21`

### Lane B: project and application service path

- `02` -> `03`, `04`, `05` -> `06`, `07`, `09`, `10`, `11`, `12`, `13`, `14`

### Lane C: predefined database path

- `15` -> `16`, `17`, `18` -> `19`, `20`

### Lane D: human decision work

- `22`
- `23`

## Best starting points for `easy-ai`

If you want the cleanest AFK start order, use:

1. `01` Install Godploy on Ubuntu from GHCR
2. `02` Create a Project with a default production Project Instance
3. `15` Create a Postgres Predefined Database Service with an Internal URL

Reason:
- They are independent.
- They each produce a demoable vertical slice.
- Together they establish the runtime, project-instance, and database-service foundations for almost everything else.

## Highest-risk slices

- `01` because installer and runtime assumptions affect every later demo.
- `21` because the GitHub manifest, webhook, and public URL contract affects all preview automation.
- `06` because preview instance creation is the pivot point for most application-service follow-up slices.
- `18` because orphan-volume lifecycle design affects both restore and project deletion behavior.

## Lowest-risk slices

- `16` because it should reuse the predefined-service template flow from `15`.
- `13` because it is mostly presentation and state modeling after preview lifecycle behavior already exists.
- `20` because it follows an already-established orphan-volume model from `18`.
