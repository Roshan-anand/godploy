# extras-v1 — 08: Predefined DB — Volume Size, Stop & Start

## What to build

Two additions to the **Predefined Database Service** lifecycle.

### Volume Size

- On creation of a Postgres/Redis predefined DB, show a "Volume Size" input field (e.g., in GB).
- A sensible default (e.g., 1GB or 10GB depending on context) — pre-filled but editable.
- In the predefined DB settings, the volume size is editable.
- On save, the backend stores the updated size. Changes take effect on next redeploy (consistent with the existing "edit requires redeploy" pattern for predefined DBs).

### Stop & Start (distinct from Delete)

- **Stop** — takes the predefined DB offline while preserving its configuration and data volume.
  - Backend: scale the swarm service to 0 replicas, keep the volume intact, mark the service status as "stopped".
  - UI: the service card shows "Stopped" instead of "Running". A "Start" button is available.
- **Start** — brings the predefined DB back online.
  - Backend: scale the swarm service back to 1 replica, the service reuses its existing volume.
  - UI: status returns to "Running".
- These are separate from "Delete" which (as before) offers optional data retention as an orphan volume.
- Stop and Start do NOT involve redeploy — the existing config is reused.

## Acceptance criteria

- [ ] Volume size field on predefined DB creation form (pre-filled default, user-editable)
- [ ] Volume size editable in predefined DB settings
- [ ] Stop action takes the DB offline, preserves config + data, shows "Stopped" status
- [ ] Start action brings the DB back online, reuses existing volume, shows "Running"
- [ ] Stop/Start distinct from Delete in both UI and backend
- [ ] Volume size changes require redeploy to take effect

## Blocked by

None — can start immediately.
