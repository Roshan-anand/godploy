# Queries

- [ ] what is embed.FS
    ```go
    //go:embed all:dist
    var embedded embed.FS

    var DistDirFS, _ = fs.Sub(embedded, "dist")
    ```

- [ ] pooling
    - what why when how pooling
    - simple example Go code for pooling

- [ ] how to production JWT
    - how to use JWT in production manner
    - what are best practices

- [ ] wht is COLESCE in SQL
    - what is COALESCE in SQL
    - how it pairs with GROUP BY

- [ ] CSRF deep dive

- [ ] AES encryption
    - what is AES encryption
    - what is AES-256-GCM

- [ ] tanstack query lazy fetch for org switcher
    - how `enabled: false` + `refetch()` works for click-to-load dropdown data
    - when to update local store from query cache vs mutation response

- [ ] svelte class based global context [context.svelte.ts](../frontend/src/lib/components/services/context.svelte.ts)

- [ ] provider integration state via HTTP status codes
    - when to use `409 Conflict` vs `204 No Content` for integration-driven UI
    - how status-code contracts simplify frontend state branches

- [ ] github app selection before repo discovery
    - how `/provider/github/app/list` can populate a client-side picker
    - how repo fetches should carry the selected `app_id`

- [ ] app service source metadata persistence
    - how `git_branch` and `build_path` should be validated in API input
    - how those fields flow from form payload to `app_service` schema/sqlc params
