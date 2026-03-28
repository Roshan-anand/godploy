# MVP RoadMap

#### Goal : User able to install and authenticate into app. create project and service (Predefined). connect Github and pull build and deploy the repo. manage env ssl and domain.

## Stories : As as user ...

- [ ] i want to install Godploy on my vps in one cmd.
- [x] i want to register/login using email and password.
- [ ] i want to CRUD organizations.
- [ ] i want to CRUD projects.
- [ ] i want to CRUD services.
- [ ] i want to setup DB services (PSQL, MongoDB).
- [ ] i want to connect to Github.
- [ ] i want to select a repo and branch to deploy.
- [ ] i want to view the build logs.
- [ ] i want to view the service logs.
- [ ] i want to set ssl certs for the service.
- [ ] i want to set custom domain for the service.

## Tasks :

- **Setup**
  - [ ] Install script (.sh) to setup Godploy and Traefik.
  - [ ] Uninstall script (.sh) to remove Godploy and Traefik.
  - [ ] Setup localtunnel for local webhook testing.
  - [ ] Setup Traefik for domain routing and ssl management.
  - [ ] Setup docker swarm mode.

- **Authentication**
  - [ ] User Registration and Login (email/password).
  - [ ] JWT-based authentication for API access.
  - [ ] Password reset functionality.
  - [ ] invite team members via email.
  - [ ] RBAC

- **ORG/project**
  - [ ] CRUD for Org.
  - [ ] CRUD for Projects.

- **Services**
  - [ ] CRUD for Services.
  - [ ] Deploy, Stop, Rebuild service.
  - [ ] Predefined service templates (PSQL, MongoDB).
  - [ ] Application Service template (for user code deployments).
  - [ ] Build logs streaming.
  - [ ] Service logs streaming.
  - [ ] Environment variable management.
  - [ ] SSL certificate management (Let's Encrypt integration).
  - [ ] Custom domain management (Traefik integration).

- **OCI Image Builder**
  - [ ] Build images using Dockerfile.
  - [ ] Build images using Nixpacks.
  - [ ] Build images using Buildpacks.

- **Github Integration**
  - [x] installation of GitHub App.
  - [ ] Fetch the repositories and branches.
  - [ ] WebHooks to trigger auto-deploy on push.
  - [ ] Build context management - Cleaning up old build artifacts.
