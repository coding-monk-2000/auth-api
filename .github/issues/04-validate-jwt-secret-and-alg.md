# Security: Validate JWT secret presence and token signing algorithm

## Description
`utils/jwt.go` reads `JWT_SECRET` at package init which may be empty. It also parses tokens without checking the signing method.

## Acceptance Criteria
- Application fails to start if `JWT_SECRET` is empty.
- Token validation explicitly checks for HMAC signing method before accepting token.
- Add unit tests for invalid/missing secret and unexpected signing algorithm.

## Implementation Notes
- Move secret loading/validation to `main` or a `config` package and inject into utils or an auth service.
- Use `jwt.ParseWithClaims` and check `token.Method`.

## Labels
`security`, `bug`, `backend`
