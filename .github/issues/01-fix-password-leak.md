# Fix: Don't return hashed password in Register response

## Description
Register currently returns the full `User` object including the hashed `password` field. This leaks sensitive data and is unnecessary.

## Acceptance Criteria
- Register endpoint response does not include the `password` field.
- Response contains only safe fields (e.g., `id`, `username`, `email`, `created_at`).
- Unit test added to assert password is not present in JSON response.

## Implementation Notes
- Use a response DTO or omit the password field (`json:"-"`) when encoding response.
- Update any clients or docs if the response shape changes.

## Labels
`bug`, `security`, `backend`
