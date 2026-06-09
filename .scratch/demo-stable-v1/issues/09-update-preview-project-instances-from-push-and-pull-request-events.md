# Update preview Project Instances from push and pull request events

Status: easy-ai

## What to build

Tighten GitHub-triggered deploy behavior so push and pull request events update only the matching services inside the matching production or preview **Project Instance**. The user-visible result should be predictable updates based on repository, selected **Git Source**, and `watch_path` matching rather than broad repo-wide rebuilds.

## Acceptance criteria

- [ ] Push and pull request update events only schedule deployments for the matching services in the matching instances.
- [ ] Repo and `watch_path` matching can expand which services inside a preview instance switch to PR or branch source when new commits touch additional watched paths.
- [ ] Services that are unrelated to the event are not rebuilt in the same or other instances.

## Blocked by

- `.scratch/demo-stable-v1/issues/07-create-a-pull-request-preview-project-instance-from-available-pr-candidates.md`
- `.scratch/demo-stable-v1/issues/21-align-github-app-manifest-webhook-endpoint-and-public-server-url-behavior.md`
