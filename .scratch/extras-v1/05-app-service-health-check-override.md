# extras-v1 — 05: App Service — Health Check Override

## What to build

Allow the user to specify a custom health check for an application service, with the ability to override any health check defined in the Dockerfile.

**Creation form:**
- Add an optional "Health Check" input field when creating an application service.
- Place a hint beneath the input: *"Leave empty to use the Dockerfile-defined health check. Enter a value to override it."*
- The health check value is stored in the DB as an optional field.

**Service settings:**
- The health check field is editable in the service settings page.
- On save, the backend reconfigures the Docker Swarm service with the new health check.

**Runtime behavior:**
- When no health check is stored, the Dockerfile's health check (if any) is used.
- When a health check is stored, it overrides any Dockerfile health check in the swarm service definition.
- The runtime status aggregation module (already built) prefers container health checks for status display.

**Health check format:**
- A simple CMD-like string or the full Docker Compose health check structure. Define one consistent format (e.g., `curl -f http://localhost:3000/health`).

## Acceptance criteria

- [ ] Optional health check field visible during service creation
- [ ] Hint text displayed below the input field
- [ ] Health check editable in service settings
- [ ] Backend stores the health check value and passes it to Docker Swarm creation/update
- [ ] Stored health check overrides any Dockerfile-defined health check at runtime
- [ ] Empty health check field means "use Dockerfile health check"

## Blocked by

None — can start immediately.
