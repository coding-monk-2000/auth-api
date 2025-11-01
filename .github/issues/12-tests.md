# Task: Add unit and integration tests

## Description
Repository lacks tests. Add unit tests for JWT utils, handlers, storage (using in-memory sqlite), and middleware.

## Acceptance Criteria
- Unit tests for `utils.GenerateToken` and `ValidateToken` covering valid, expired, and malformed tokens.
- Handler tests for Register/Login using a test store or sqlite in-memory DB.
- CI runs tests on push/PR.

## Implementation Notes
- Use `testing` package and table-driven tests. Consider `httptest` for handler tests.

## Labels
`tests`, `backend`
