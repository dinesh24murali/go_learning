# Gin learn

I need to find suitable alternatives for the following from `NodeJS`

1. Middlewares at route level *
2. CORS *
3. Express router * - `r.Group("/api")`
4. Read environment variables *
5. logger for go *
6. Middleware at global level. Something equivalent to `app.use`. *
7. Something equivalent to `yarn workspaces`
8. Schema validator `joy`, `zod` *
9. Swagger

Run the server and watch for file changes:
```bash
find . -name "*.go" | entr -r go run .
```