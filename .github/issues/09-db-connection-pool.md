# Improvement: Configure DB connection pool and optional migrations

## Description
Database initialization does not configure connection pool settings. For production, pool tuning and controlled migrations are required.

## Acceptance Criteria
- Configure `sql.DB` pool settings (MaxOpenConns, MaxIdleConns, ConnMaxLifetime) based on env vars.
- Add a toggle to run AutoMigrate via a `MIGRATE_ON_START` env var (default false for safety).
- Document recommended values in README.

## Implementation Notes
- Use `db.DB()` to obtain `*sql.DB` and set options.

## Labels
`enhancement`, `database`, `ops`
