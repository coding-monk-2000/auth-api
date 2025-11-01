# Task: Add CI for linting and tests

## Description
Add GitHub Actions workflow to run `gofmt`, `go vet`, `golangci-lint` (or staticcheck), and `go test` on PRs.

## Acceptance Criteria
- A `.github/workflows/ci.yml` that runs on push/PR and fails on formatting/lint/test failures.
- Documentation in README on how to run linters locally.

## Labels
`ci`, `devops`
