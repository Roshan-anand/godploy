# TODOs

## staging

- [ ] setup openAPI
- [ ] Setup logger
- [x] setup health check route
- [x] setup tests
- [ ] integration test for main func
- [ ] setup CI/CD
- [ ] setup linting and formatting
- [x] write documentation
- [ ] setup monorepo
- [x] setup hot reloading in development
- [x] setup dev env for separate frontend
- [x] setup containerized dev env
- [x] migrate UI from vue to svelte

## application

- [ ] setup rate limiting middleware [medium, backend]
- [ ] delete project btn sould show popup to confirm deletion. [easy, ui]
- [ ] git provider page. [easy, ui]
- [ ] tests for project lifecycle
- [ ] tests for service lifecycle
- [ ] rememberMe functionality for login

## enhancements

- [ ] use dynamic imports in client side

## Potential bugs

- [x] not using enums for column user.role in [sql](../sqlite/migrations/0001_init_schema.up.sql). **[ solved ]**
- [ ] service data is stored in DB and deployed, but what if user remove the service from terminal. the data still exists.
- [ ] the github app is stored linked to org_ig, what if user fails on instllation then data still ramins. so retry fails because there is multiple github app store in singe org.
