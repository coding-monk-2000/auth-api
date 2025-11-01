# Fix: Handle environment and DB initialization errors on startup

## Description
`main.go` currently ignores errors from `godotenv.Load()` and `storage.InitDatabase()`. If the DB or required env vars are missing the server may start in a broken state.

## Acceptance Criteria
- Application fails fast when required env variables (JWT_SECRET, DB_DRIVER, etc.) are missing.
- `InitDatabase` errors are handled and cause process exit with a clear log message.
- Add tests or a runbook describing required env vars.

## Implementation Notes
- Return/propagate errors from `InitDatabase` and call `log.Fatal` or exit in `main` when errors occur.

## Labels
`bug`, `reliability`, `backend`
