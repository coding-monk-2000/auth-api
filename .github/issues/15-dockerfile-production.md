# Improvement: Harden Dockerfile for production

## Description
Dockerfile should be multi-stage, use a small runtime image, and set non-root user and proper env var defaults.

## Acceptance Criteria
- Multi-stage build that compiles binary in `golang:alpine` (or similar) and copies into `scratch`/`distroless`/`gcr.io/distroless/base` image.
- Non-root user and minimal layers.
- Document build/run process in README.

## Labels
`ops`, `docker`
