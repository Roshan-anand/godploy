# Align GitHub App manifest, webhook endpoint, and public server URL behavior

Status: hard-human

## What to build

Turn the open GitHub integration contract into a concrete V1 decision and implementation path. The result should align the GitHub App manifest, webhook endpoint, and public server URL behavior so branch and pull request automation has one clear runtime contract.

## Acceptance criteria

- [ ] The intended public server URL and callback or webhook contract for GitHub is explicitly decided.
- [ ] The GitHub App manifest and webhook behavior match that decision rather than relying on implicit or conflicting assumptions.
- [ ] The resulting contract is stable enough that later preview and webhook slices can build on it without re-deciding the integration shape.

## Blocked by

- `.scratch/demo-stable-v1/issues/01-install-godploy-on-ubuntu-from-ghcr.md`
