# Refactor: Consolidate configuration into a config package/struct

## Description
Environment variables are read from multiple places. A central `config` package will validate and provide typed config to components.

## Acceptance Criteria
- Add `config` package that reads env vars and returns a typed struct (Port, DB config, JWT secret, timeouts).
- `main` uses the config and fails fast on missing required values.
- Replace ad-hoc env reads across the codebase.

## Implementation Notes
- Provide a `NewConfigFromEnv()` function with validation.

## Labels
`refactor`, `backend`
