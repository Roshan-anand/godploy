# Demo-Stable V1 PRD

Source PRD: `docs/prd.md`

This tracker directory holds the vertical-slice issues derived from the current V1 PRD for Godploy's demo-stable release.

## Scope summary

- Install Godploy on Ubuntu from GHCR
- Normalize the product model around `Organization -> Project -> Project Instance -> Service`
- Create an explicit production `Project Instance` for every `Project`
- Support isolated branch and pull request preview `Project Instances`
- Keep GitHub-driven preview automation aligned with `repo_id` and `watch_path` rules
- Ship `Predefined Database Service` flows for Postgres and Redis inside instance-scoped runtime
- Preserve and reattach `Orphan Volume` data through `Storage`
- Tighten runtime visibility, instance lifecycle visibility, and remaining V1 operational decisions
