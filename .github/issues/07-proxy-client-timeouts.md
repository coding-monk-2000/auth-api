# Improvement: Add timeouts and header proxying to ProxyToMessages

## Description
`handlers.ProxyToMessages` uses a default http.Client without timeouts and doesn't copy response headers back to the client.

## Acceptance Criteria
- Use an `http.Client` with a reasonable timeout (e.g., 5s).
- Copy relevant headers (Content-Type, Cache-Control, etc.) from the proxied response back to the client.
- Return 502 for upstream unreachable and proper error messages.
- Add integration test or lightweight e2e test for proxy behavior.

## Labels
`enhancement`, `backend`, `ops`
