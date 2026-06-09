# Warn about Orphan Volumes when deleting a Project

Status: easy-ai

## What to build

Add the project-delete warning path so users understand what preserved database data is attached to the project they are removing. The delete flow should surface related **Orphan Volumes** and make the preservation or removal outcome explicit.

## Acceptance criteria

- [ ] Deleting a **Project** warns the user when related **Orphan Volumes** exist.
- [ ] The flow makes the preservation or removal consequence explicit before the project delete completes.
- [ ] Any preserved orphaned data follows the intended post-delete ownership behavior instead of becoming hidden state.

## Blocked by

- `.scratch/demo-stable-v1/issues/18-preserve-deleted-database-data-as-an-orphan-volume-and-show-it-in-storage.md`
