# Install Godploy on Ubuntu from GHCR

Status: easy-ai

## What to build

Deliver the first-run install path for demo-stable V1 on Ubuntu. A user should be able to run the installer, get Docker and swarm bootstrapped if needed, pull the Godploy and Traefik images from GHCR, provision persistence, and reach the dashboard on the documented address.

## Acceptance criteria

- [ ] A fresh Ubuntu VPS can install Godploy from GHCR through the documented `install.sh` flow.
- [ ] The install path provisions the runtime pieces required for Godploy, Traefik, and persisted metadata storage.
- [ ] A user can reach the Godploy dashboard successfully after install without doing undocumented manual setup.

## Blocked by

None - can start immediately
