# Support manual delete and TTL cleanup for branch preview Project Instances

Status: easy-ai

## What to build

Finish the branch-preview lifecycle by supporting both manual deletion and TTL-driven cleanup for branch preview **Project Instances**. Users should be able to reclaim temporary preview runtimes explicitly, and Godploy should also be able to clean them up automatically when a configured lifetime expires.

## Acceptance criteria

- [ ] A user can delete a branch preview **Project Instance** manually from the dashboard.
- [ ] Branch preview instances can be configured or recorded for TTL-based async cleanup.
- [ ] Manual and TTL cleanup both remove the full cloned preview runtime rather than only part of the preview resources.

## Blocked by

- `.scratch/demo-stable-v1/issues/06-create-a-branch-preview-project-instance-from-a-production-snapshot.md`
