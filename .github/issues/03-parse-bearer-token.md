# Improvement: Parse `Authorization: Bearer <token>` header correctly

## Description
Handlers and `utils.ValidateToken` currently expect the raw token string. In practice clients send `Authorization: Bearer <token>`.

## Acceptance Criteria
- All places that read `Authorization` header accept and correctly handle the `Bearer ` prefix.
- Add a small helper to extract the token and reuse it from middleware.
- Unit tests for parsing behavior.

## Implementation Notes
- Use `strings.HasPrefix(strings.ToLower(authHeader), "bearer ")` and trim the prefix.

## Labels
`enhancement`, `backend`
