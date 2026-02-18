# Godploy — Product Requirements Document

## What is Godploy?

Godploy is a lightweight, single-binary, self-hosted PaaS (Platform as a Service). Think of it as a simpler alternative to tools like Dokploy and Coolify.

You run one binary on your server — it gives you a web dashboard where you can connect your Git repos, build container images, deploy services, and manage routing + TLS. Everything runs on Docker under the hood, with Traefik handling ingress.

## Why?

Self-hosting applications today means either:

- Paying for managed platforms (Vercel, Railway, Render, etc.)
- Wrangling complex tools that try to do too much

Godploy aims to be the middle ground — a minimal, easy-to-run tool that handles the core deployment loop: **push code → build image → run container → route traffic**.

## Tech Stack

| Layer       | Technology                          |
| ----------- | ----------------------------------- |
| Server      | Go (Echo framework)                 |
| Frontend    | Vue 3 SPA (embedded into the binary)|
| Database    | SQLite (metadata only, via sqlc)    |
| Runtime     | Docker (container lifecycle)        |
| Ingress     | Traefik (routing, TLS, subdomains)  |
| Auth        | JWT + session-based                 |

The Vue frontend is compiled and embedded into the Go binary at build time — so the final output is a single executable with no external dependencies beyond Docker and Traefik.

## Architecture at a Glance

```
User Browser
    ↓
[ Traefik ] ← handles TLS, domain routing
    ↓
[ Godploy Binary ]
  ├── Vue SPA (embedded static files)
  ├── REST API (Echo)
  ├── SQLite (users, orgs, projects, services, sessions)
  └── Docker Client (build, run, manage containers)
    ↓
[ Docker Engine ]
  ├── App containers (user services)
  ├── Predefined services (Postgres, MongoDB, etc.)
  └── Isolated networks per project
```

## MVP Scope

The MVP focuses on getting a working deployment pipeline end-to-end. Here are the 7 core areas:

### 1. Traefik Ingress
Setup Traefik as the main entrypoint — handle subdomain/domain routing, automatic TLS via Let's Encrypt, and route traffic to the right containers.

### 2. Git Provider Integration
OAuth-based connection to GitHub, GitLab, or Bitbucket. Fetch repos/branches, set up webhooks for auto-deploy on push.

### 3. OCI Image Building
Build container images from user repos using Dockerfile, Nixpacks, or Buildpacks. Includes build logs, queue management, and predefined service templates (Postgres, MongoDB, etc.).

### 4. Container Management
Full container lifecycle — start, stop, restart, remove. Manage Docker networks (isolated per project), volumes, secrets, resource limits, and real-time log streaming.

### 5. Monitoring & Logging
Container health checks, resource usage metrics (CPU, memory, network), and deployment history tracking.

### 6. User Authorization
Organizations, projects, teams, and role-based access control (RBAC).

### 7. Installation
Shell scripts to install/setup Godploy + Traefik on a fresh server, and to cleanly uninstall everything.

## User Stories

- Login to the dashboard
- Create / delete / view projects
- Create predefined services (Postgres, MongoDB)
- Connect a Git provider and select a repository
- Build and deploy a service from a repo
- View logs and status of running services
- Create / delete custom services

## Current Status

- Auth system is implemented (register, login, JWT + sessions)
- SQLite schema covers users, orgs, projects, services, sessions
- Basic API structure with Echo
- Traefik config is in staging
- Vue frontend is scaffolded

---

> For the detailed feature checklist and task tracking, see [`mvp_roadmap.md`](./mvp_roadmap.md) and [`todos.md`](./todos.md).
