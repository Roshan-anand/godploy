# Create a Project with a default production Project Instance

Status: easy-ai

## What to build

Extend project creation so every new **Project** automatically gets one explicit production **Project Instance**. The user-visible result should be that a project is no longer just a container for services; it is also the owner of the primary runtime instance that later preview instances clone from.

## Acceptance criteria

- [x] Creating a **Project** also creates one production **Project Instance** automatically.
- [x] The production **Project Instance** is stored and retrievable as explicit project state rather than implied runtime behavior.
- [x] The dashboard and API can show that a newly created project has a production instance even before any services are added.

## Blocked by

None - can start immediately
