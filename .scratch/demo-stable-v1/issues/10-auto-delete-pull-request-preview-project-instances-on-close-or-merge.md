# Auto-delete pull request preview Project Instances on close or merge

Status: easy-ai

## What to build

Complete the pull request preview lifecycle so PR preview **Project Instances** are cleaned up automatically when their pull request is closed or merged. Cleanup should happen asynchronously and remove the cloned runtime instead of leaving stale previews behind.

## Acceptance criteria

- [ ] Pull request close and merge events enqueue deletion of the matching preview **Project Instance** automatically.
- [ ] Preview cleanup removes the cloned services, deployments, volumes, and network instead of leaving partially active runtime state.
- [ ] Cleanup failures leave the preview in a recoverable lifecycle state such as deleting or error rather than disappearing silently.

## Blocked by

- `.scratch/demo-stable-v1/issues/07-create-a-pull-request-preview-project-instance-from-available-pr-candidates.md`
- `.scratch/demo-stable-v1/issues/08-keep-open-pull-request-candidates-in-sqlite-and-show-them-in-the-dashboard.md`
