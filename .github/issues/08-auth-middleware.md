# Feature: Implement authentication middleware

## Description
Token parsing/validation logic is duplicated in handlers. Implement middleware to validate JWT, attach claims to request context, and reuse across routes.

## Acceptance Criteria
- New middleware `AuthMiddleware` that checks Authorization header and validates JWT.
- Middleware stores parsed claims (username) in request context for handlers to access.
- Replace ad-hoc validation in `ProxyToMessages` and expose a route-protected example.
- Unit tests for middleware behavior.

## Implementation Notes
- Keep middleware small and fast; return 401 on validation failure.

## Labels
`feature`, `backend`
