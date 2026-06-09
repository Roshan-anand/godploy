# Show project instance switching in the dashboard

Status: easy-ai

## What to build

Update the project dashboard so the user first selects a **Project Instance** and then sees the services inside that selected runtime. The result should make the production-versus-preview boundary obvious throughout the normal project UI.

## Acceptance criteria

- [x] The project dashboard shows available **Project Instances** for the selected **Project**.
- [ ] Switching the selected instance updates the service list and detail navigation to that instance's services.
- [ ] The UI makes it clear whether the user is viewing the production instance or a preview instance.

## Blocked by

- `.scratch/demo-stable-v1/issues/02-create-a-project-with-a-default-production-project-instance.md`