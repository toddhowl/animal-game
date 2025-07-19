# Animal Guessing Game in Go

This is a simple command-line game written in Go by my son. The program randomly selects an animal from a data file and gives you hints about it. Your goal is to guess the animal based on the hints provided.

## Features
- Randomly selects an animal from a list in `animals.json`
- Provides multiple hints for each animal
- Accepts user guesses and tells you if you are correct (supports alternate/short names)
- All code is well-commented and formatted for learning purposes

## How to Run

1. **Install Go** (if you don't have it):
   - Download and install from [https://golang.org/dl/](https://golang.org/dl/)

2. **Clone or Download this Repository**

3. **Run the Program**
   - Open a terminal in the project directory
   - Run:
     ```sh
     go run main.go
     ```

## How to Play
- The program will give you a series of hints about a randomly chosen animal.
- After each hint, you can type your guess.
- If you guess correctly, you win!
- If not, you get another hint (if available).
- After all hints are used, the answer is revealed.

## Customizing Animals and Hints
- All animal data is stored in `animals.json` as an array of objects.
- To add or edit animals, open `animals.json` and add a new object with fields:
  - `name`: the main animal name (string)
  - `altNames`: (optional) array of alternate/short names
  - `hints`: array of hint strings
- Keep the JSON valid (no comments, no trailing commas). Use a JSON formatter for large edits.

## Example
```
Hint 1: I am the largest land animal.
Your guess: lion
Incorrect!
Hint 2: I have a trunk.
Your guess: elephant
Correct! The animal was: elephant
```

## Example animal entry in animals.json
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

## License
This project is open source and free to use for learning purposes.
