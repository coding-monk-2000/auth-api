# Fix: Handle DB "record not found" and other GORM errors properly

## Description
Handlers assume `GetUser` returns `nil` when not found. GORM returns `ErrRecordNotFound` which should be handled explicitly.

## Acceptance Criteria
- `GetUser`/handlers distinguish between "not found" (return 401 on login) and real DB errors (500).
- No nil-pointer or misleading error messages.
- Unit tests for both scenarios.

## Implementation Notes
- Use `errors.Is(err, gorm.ErrRecordNotFound)` in handlers.

## Labels
`bug`, `backend`
