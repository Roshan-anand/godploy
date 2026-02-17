# MVP RoadMap

## Goal 
 A working MVP of Godploy. It is a selfhost PAAS alternative to Dokploy, coolify. here users can come connect to any git Provider, and deploy their projects by definfin their repo. 

## Core aspects

- **Traefik**
  - [ ] Setup Traefik as ingress if not installed.
  - [ ] Setup Traefik as main entrypoint for all the services.
  - [ ] SSL/TLS certificate management - Let's Encrypt integration or custom certs.
  - [ ] Subdomain/domain routing - Map services to domains automatically.
- **Git Provider**
  - [ ] Connect to any Git provider (Github, Gitlab, Bitbucket).
  - [ ] Fetch the repositories and branches.
  - [ ] WebHooks to trigger auto-deploy on push.
  - [ ] OAuth flow for secure repository access.
  - [ ] Build context management - Cleaning up old build artifacts.
- **OCI Image Builder**
  - [ ] Build images using Dockerfile.
  - [ ] Build images using Nixpacks.
  - [ ] Build images using Buildpacks.
  - [ ] Build Logs, queue management for multiple builds.
  - [ ] Predefined Services (PSQL, MongoDB) - Template-based image building.
- **Container Management**
  - [ ] Manage Container Lifecycle (start, stop, restart, remove).
  - [ ] Manage Docker Networks (create isolated networks per project).
  - [ ] Manage Volumes (persistent data for databases and user services).
  - [ ] Stream Container Logs (real-time log access via WebSocket/SSE).
  - [ ] Secrets Management.
  - [ ] Resource Limits (CPU/memory constraints for services).
  - [ ] Rollback capability - Revert to previous deployments.
- **Monitoring & Logging**
  - [ ] Container health checks - Detect crashed services.
  - [ ] Resource usage metrics - CPU, memory, network per service.
  - [ ] Deployment history - Track what was deployed when.
- **User Athorization**
  - [ ] manage Organizations and Projects access.
  - [ ] Role-based access control (RBAC) for users within organizations.
  - [ ] manage whole teams.
- **Installation and deletion**
  - [ ] Install script (.sh) to setup Godploy and Traefik.
  - [ ] Uninstall script (.sh) to remove Godploy and Traefik.

## User Stories

- [ ] able to login to Dashboard
- [ ] able to create a new project.
- [ ] able to delete a project.
- [ ] able to view all the projects in the dashboard.
- [ ] able to create a new pre-defined service(PSQL, MongoDB).
- [ ] able to connect to a Git Provider and select a repository.
- [ ] able to build and deploy the service.
- [ ] able to view the logs of the service.
- [ ] able to create a new custom service.
- [ ] able to view all the services in the dashboard.
- [ ] able to view the logs, status of the services .
- [ ] able to delete a service.
