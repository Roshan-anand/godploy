# Create a branch preview Project Instance from a production snapshot

Status: easy-ai

## What to build

Deliver the end-to-end flow for creating a preview **Project Instance** from a manually selected branch. Creating the preview should snapshot the current production instance, clone all services into the preview, give the preview its own **Instance Network**, rebuild only the services selected by the branch rules, reuse pinned ready images for unchanged application services, and provision fresh isolated stateful services.

## Acceptance criteria

- [ ] A user can create a preview **Project Instance** from any selected branch even when no webhook record exists for that branch.
- [ ] Preview creation clones the full production service topology into a separate instance with its own private network and fresh stateful service data.
- [ ] Only the selected branch-targeted services rebuild from source while unchanged application services reuse pinned ready production images.

## Blocked by

- `.scratch/demo-stable-v1/issues/03-create-an-internal-application-service-inside-the-production-project-instance.md`
- `.scratch/demo-stable-v1/issues/05-show-project-instance-switching-in-the-dashboard.md`
