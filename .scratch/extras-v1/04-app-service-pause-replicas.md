# extras-v1 — 04: App Service — Pause & Replicas

## What to build

Give the user control over an application service's runtime presence and capacity.

**Pause / Resume:**
- A "Pause" action button in the service settings or detail page.
- Pausing sets the swarm service replicas to 0 — container stops, config preserved.
- A disabled/paused indicator shows on the service card and detail page.
- Resume restores replicas to the previous count (or a default of 1).

**Replicas:**
- In the service settings, show current replica count with increment (+) and decrement (−) buttons.
- Minimum 0 (paused), no hard maximum but a sensible upper bound warning.
- Backend updates the swarm service replica count.

Both operations call a backend endpoint that updates the Docker Swarm service scale. The updated count is persisted in the DB and reflected in the UI immediately after the operation completes.

## Acceptance criteria

- [x] Pause sets replicas to 0; service shows as "paused" in the UI
- [x] Resume restores replicas to previous count; service shows as "running"
- [x] Increment/decrement button adjusts replicas and updates the UI
- [x] Backend scales the Docker Swarm service accordingly

## Blocked by

None — can start immediately.
