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

- [ ] feature-first query and mutation modules
    - keep shared query keys, fetchers, mutations, and domain types in `frontend/src/lib/features/<feature>/`
    - let route files stay focused on UI wiring and local form state
