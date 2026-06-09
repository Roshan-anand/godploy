# Create a pull request preview Project Instance from available PR candidates

Status: easy-ai

## What to build

Deliver the full dashboard flow for creating a preview **Project Instance** from an available pull request candidate. The user should be able to pick an open PR, review the affected services implied by repo and watch-path matching, create the preview, and get one active preview instance for that project and pull request.

## Acceptance criteria

- [ ] A user can create a preview **Project Instance** from an available open pull request candidate.
- [ ] Preview creation uses the PR diff and watch-path rules to decide which services switch to PR source while the rest stay pinned to the production snapshot.
- [ ] Godploy prevents duplicate active preview instances for the same project and pull request.

## Blocked by

- `.scratch/demo-stable-v1/issues/06-create-a-branch-preview-project-instance-from-a-production-snapshot.md`
- `.scratch/demo-stable-v1/issues/08-keep-open-pull-request-candidates-in-sqlite-and-show-them-in-the-dashboard.md`
