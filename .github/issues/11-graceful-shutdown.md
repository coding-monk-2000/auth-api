# Improvement: Implement graceful shutdown for HTTP server

## Description
`main` uses `http.ListenAndServe` directly. Add a graceful shutdown with context and timeout to allow in-flight requests to finish.

## Acceptance Criteria
- Replace `ListenAndServe` with `http.Server` and `server.ListenAndServe()` in a goroutine.
- Catch SIGINT/SIGTERM and call `server.Shutdown(ctx)` with a configurable timeout.
- Log shutdown steps and return non-zero exit code if shutdown fails.

## Labels
`enhancement`, `reliability`
