# Improvement: Validate request bodies for Register and Login

## Description
API currently decodes JSON without validating required fields (username, password, email) and accepts empty values.

## Acceptance Criteria
- Validate incoming JSON for required fields and format (e.g., email regex if provided).
- Return 400 with a structured error message listing missing/invalid fields.
- Add unit tests covering validation failures.

## Implementation Notes
- Consider using a small validation helper or library (e.g., go-playground/validator).

## Labels
`enhancement`, `backend`
