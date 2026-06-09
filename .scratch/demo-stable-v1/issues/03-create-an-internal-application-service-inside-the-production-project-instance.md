# Create an internal application Service inside the production Project Instance

Status: easy-ai

## What to build

Deliver the full logged-in flow for creating an internal application **Service** inside a production **Project Instance**. The flow should let the user choose the project, provider connection, repository, initial branch, and **Exposure Mode**, then create a service that joins the production instance's **Instance Network** without receiving public ingress.

## Acceptance criteria

- [x] The create-service flow persists the application **Service** under the production **Project Instance** of the selected **Project**.
- [x] The form lets the user choose the initial repository branch so the first production **Git Source** is explicit.
- [x] Choosing internal access results in a service that is reachable through the production instance network but not exposed publicly.

## Blocked by

- `.scratch/demo-stable-v1/issues/02-create-a-project-with-a-default-production-project-instance.md`
