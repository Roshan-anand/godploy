# All API Endpoints

- [x] GET /api/auth/user :
  - validate JWT -> isAuth goto /dshboard
  - else check accExists goto /login else /register

- [x] POST /api/login : validate pass, generate JWT -> goto /dashboard

- [x] POST /api/register : create acc, generate JWT -> goto /dashboard

- [ ] GET /api/orgs : return all organisations

- [ ] GET /api/projects?orgId= : returns all projects of an org

- [ ] GET /api/project : return all service info of the project

- [ ] POST /api/project : create a new project

- [ ] DELETE /api/project :
  - if noService running remove project
  - else warn

- [ ] GET /api/service : return info of a service

- [ ] POST /api/service : create new empty service

- [ ] DELETE /api/service : remove and stop the service

