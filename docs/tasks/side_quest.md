# TODOs

## staging

- [ ] setup openAPI
- [ ] Setup logger
- [ ] setup health check route
- [x] setup tests
- [ ] integration test for main func
- [ ] setup CI/CD
- [ ] setup linting and formatting
- [x] write documentation
- [ ] setup monorepo
- [x] setup hot reloading in development
- [x] setup dev env for separate frontend
- [ ] setup containerized dev env

## application
- [ ] rememberMe functionality for login
- [ ] setup rate limiting middleware
- [ ] tests for project lifecycle
- [ ] tests for service lifecycle

## enhancements
- [ ] use dynamic imports in client side

## Potential bugs

- [x] not using enums for column user.role in [sql](../sqlite/migrations/0001_init_schema.up.sql). **[ solved ]**
- [ ] service data is stored in DB and deployed, but what if user remove the service from terminal. the data still exists.