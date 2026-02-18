# Queries

- [ ] AI Knowledge Capture
    - When AI introduces new design/concept in code, add summary comment explaining the reasoning
    - Update `/docs/queries.md` with the topic and brief explanation for future reference
    - This creates a searchable knowledge base of engineering decisions

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