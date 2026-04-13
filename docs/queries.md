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

- [ ] provider setup lifecycle cleanup
    - how to model provider status as nullable query response (`null` means not connected)
    - admin-only cleanup endpoint design for deleting provider credentials safely

- [ ] discriminated unions for mixed service APIs
    - how frontend uses `type` (`psql | app`) to safely render one details page with different fields
    - how to pass route context (query params vs path params) when list items share one UI but call different endpoints
