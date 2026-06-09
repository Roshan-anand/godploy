# Keep open pull request candidates in SQLite and show them in the dashboard

Status: easy-ai

## What to build

Add the PR-candidate intake path for preview creation. GitHub webhook events should add and remove open pull request metadata in SQLite, and the dashboard should surface those open PRs as selectable preview candidates for the matching repositories.

## Acceptance criteria

- [ ] Pull request open and reopen events add or refresh the matching open PR metadata in SQLite.
- [ ] Pull request close and merge events remove the matching open PR metadata from SQLite.
- [ ] The dashboard can show available open pull request candidates for preview creation without requiring the user to leave the app.

## Blocked by

- `.scratch/demo-stable-v1/issues/21-align-github-app-manifest-webhook-endpoint-and-public-server-url-behavior.md`
