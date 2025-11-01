# Feature: Add health and readiness endpoints

## Description
Add `/healthz` and `/readyz` endpoints for orchestration and readiness checks.

## Acceptance Criteria
- `/healthz` returns 200 if service process is alive.
- `/readyz` returns 200 only if DB connection (or other critical dependencies) is available.
- Add simple tests and document endpoints in README.

## Labels
`feature`, `ops`
