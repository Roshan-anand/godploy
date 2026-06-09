# extras-v1 — 09: Orphan Volume — Filters & Rename

## What to build

Enhance the existing **Storage** page (orphan volume list) with filtering and rename capability.

### Filters

Three filter controls above the volume list:

1. **Size — ascending/descending** — toggle button/icon to sort by volume size. Ascending by default. Arrow indicator shows current direction.
2. **Service type** — dropdown to filter by originating service type: "All", "Postgres", "Redis". Selecting a type shows only orphan volumes that came from that predefined DB type.
3. **Name** — a text input that filters volumes whose name contains the search string (case-insensitive). Live filtering as the user types.

Filters compose: e.g., name "staging" + type "Postgres" + size descending shows all Postgres volumes with "staging" in the name, sorted largest first.

### Rename

Each orphan volume row has a rename action (pencil icon or inline edit). Clicking opens an inline edit or a small modal to change the volume's name. The name is a display label only — it doesn't affect the underlying data or volume identifiers.

Backend: PATCH endpoint to update the orphan volume name. The existing list endpoint should accept filter query params (`sort_by`, `sort_dir`, `service_type`, `name_search`).

## Acceptance criteria

- [ ] Size sort toggle (asc/desc) works, ascending by default
- [ ] Service type filter dropdown shows "All", "Postgres", "Redis"
- [ ] Name search input filters live as the user types
- [ ] All three filters compose together correctly
- [ ] Each orphan volume row has a rename action
- [ ] Rename updates the display name in the backend and persists
- [ ] Backend accepts sort and filter query params on the list endpoint

## Blocked by

None — can start immediately.
