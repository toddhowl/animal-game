# Copilot Instructions for Animal Guessing Game (Go)

## Project Overview
- This is a Go command-line game that selects a random animal and provides hints for the user to guess the animal.
- Animal data is stored in `animals.json` as an array of objects with fields: `name` (string), optional `altNames` (array of strings), and `hints` (array of strings).
- The main logic is in `main.go`, which loads and parses `animals.json` at runtime.

## Key Files
- `main.go`: Game logic, user interaction, JSON loading, and answer checking (supports alternate names).
- `animals.json`: Animal data. Each entry must be pretty-printed, with fields in the order: `name`, `altNames` (if present), `hints`.
- `README.md`: Usage, setup, and customization instructions.

## Data & Patterns
- All animal data is externalized in `animals.json` for easy editing and expansion.
- Alternate/short names for animals are supported via the `altNames` field.
- Hints are provided as an ordered array; the game reveals them one by one.
- JSON must be valid (no comments, no trailing commas). Use a formatter for large edits.

## Developer Workflows
- **Run the game:** `go run main.go`
- **Edit animal data:** Update `animals.json` directly. Use a JSON formatter for consistency.
- **Add new animals:** Add a new object with `name`, optional `altNames`, and `hints`.
- **Debugging:** The game prints errors if `animals.json` is missing or malformed.
- **No tests**: There is no automated test suite; manual play is the primary validation.

## Project Conventions
- All code is well-commented for learning and maintainability.
- Keep `animals.json` entries consistently formatted and alphabetized if possible.
- Do not add comments or extra fields to `animals.json`.
- Use Go standard library only; no external dependencies.

## Example animal entry
```json
{
  "name": "elephant",
  "altNames": ["african elephant", "asian elephant"],
  "hints": [
    "I am the largest land animal.",
    "I have a trunk.",
    "I have tusks and big ears."
  ]
}
```

## Integration Points
- No network or external service dependencies.
- All data and logic are local to the project directory.

---

If you add new features (e.g., scoring, categories), update this file with new conventions and patterns.
