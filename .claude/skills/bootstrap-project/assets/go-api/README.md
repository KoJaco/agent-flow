# Go Lambda API skeleton (reference)

- `cmd/lambda` — Lambda entry (arm64, `provided.al2023`, build tag `lambda.norpc`).
- `cmd/local` — local dev server.
- `internal/auth` — provider-neutral JWKS verifier + middleware + origin-secret check.
- `internal/app` — wiring (chi router, pgx pool).

Setup: `go mod init <module>` (replace `example.com/app`), `go get` the imports, **pin versions**.
Env: `DATABASE_URL` (pooled), `JWKS_URL`, `JWT_ISSUER`, `JWT_AUDIENCE`, `ORIGIN_SECRET`.
This is a starting reference — compile it and adjust APIs to the versions you pin.
