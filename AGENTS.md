# Repository Guidelines

## Project Structure & Module Organization
`cmd/main.go` is the lone entry point and should stay focused on HTTP wiring. Shared configuration lives in `config/config.go`, while `internal/repository/postgres` owns pgx pools; keep any new adapters in the same tree. Schema changes remain in `migrations/`, and tests sit beside the code they cover.

## Build, Test, and Development Commands
- `SERVER_PORT=8081 go run ./cmd`: start the API with custom ports or env overrides like `DB_HOST`.
- `go build ./cmd && go vet ./...`: compile and run lightweight static checks before committing.
- `go test ./...`: execute all package tests; add `-count=1` for flaky-prone logic.
- `DB_URL='postgres://user:pass@localhost:5432/postgres?sslmode=disable' make migrate-up`: apply pending database migrations.

## Coding Style & Naming Conventions
Always run `gofmt`/`goimports` (tabs, standard spacing); do not hand-edit indentation. Exported symbols carry PascalCase names and doc comments, helpers stay camelCase, and function signatures accept `ctx context.Context` first. Keep handlers thin, avoid global mutable state, and keep migration files in the sequential `NNNNNN_description.(up|down).sql` pattern.

## Testing Guidelines
Write table-driven tests with the standard `testing` package and store them in `_test.go` files next to their targets (for example `internal/repository/postgres/db_test.go`). Favor fakes or mocked pgx pools to keep unit tests deterministic, and reserve live database checks for dedicated integration runs. Run `go test ./...` locally and cover both success and error flows for repositories and config loaders.

## Database Migrations
The `makefile` wraps the `migrate` CLI. Export a DSN-compatible `DB_URL` (matching `config.DatabaseConfig.DSN()`), then run `make migrate-up` to apply changes, `make migrate-down` for a single rollback, `make migrate-status` to confirm position, and `make migrate-create name=add_users_table` to scaffold paired SQL files. Ship migrations in the same pull request as the code that needs them.

## Commit & Pull Request Guidelines
History shows short Title Case summaries (`Init first step base Template for new project`), so keep single-line subjects under 72 characters and use bodies only when nuance is needed. Pull requests must explain what changed, why, linked issues, required env or migration steps, and include output from `go test ./...` plus any `make migrate-*` commands. Add screenshots or `curl localhost:8080/health` snippets when touching HTTP behavior.

## Security & Configuration Tips
`config.Load()` sources `SERVER_PORT`, `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, and `DB_NAME`; never embed credentials in Go code or migration files. Store secrets in your shell profile or `.envrc`, rotate `DB_PASS`, terminate TLS before requests reach this binary, and limit the database role to the privileges its migrations require.
