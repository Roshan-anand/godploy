# extras-v1 — 06: Service Domain Settings

## What to build

A settings section on each service's page to manage its domain. Behavior differs slightly between production and preview services.

**Production services:**
- Show a domain input field.
- If the service is public and a domain has been set, show it.
- If no domain is set, show an empty input with a placeholder.
- User can enter any domain. Backend stores it and applies it to Traefik routing.

**Preview services:**
- On creation, the backend auto-generates a domain following the `<service>.<preview>.<base_domain>` pattern.
- The domain section shows the auto-generated domain with two options:
  1. **Auto-generate from production** — a button that triggers re-generation from the production service's domain pattern (doesn't copy the prod domain value, just regenerates the preview pattern).
  2. **Custom domain** — an input field to enter any domain, which replaces the auto-generated domain.
- The domain section clearly labels which mode is active: "Auto-generated" or "Custom".

**Backend:**
- A service domain field (nullable) on the service record.
- Auto-generation logic for previews (runs at preview creation and optionally on-demand).
- Custom domain overwrites the auto-generated value.
- Traefik label generation uses the stored domain value.

## Acceptance criteria

- [ ] Production service settings show domain input; user can set or change it
- [ ] Preview service settings show auto-generated domain with badge/label
- [ ] "Auto-generate from production" button regenerates the preview domain pattern
- [ ] Custom domain input replaces auto-generated domain for that preview service
- [ ] Backend stores domain value and Traefik respects it
- [ ] Preview domain shows mode indicator: "Auto-generated" or "Custom"

## Blocked by

None — can start immediately.
