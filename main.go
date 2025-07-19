// Animal Guessing Game
// --------------------
// This program selects a random animal from a JSON file and gives the user hints to guess it.
// Animal data is stored in animals.json, supporting alternate names and multiple hints.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Animal represents a single animal entry loaded from animals.json.
// - Name: the canonical name of the animal
// - AltNames: alternate/short names (optional)
// - Hints: ordered list of hints for guessing
type Animal struct {
	Name     string   `json:"name"`
	AltNames []string `json:"altNames,omitempty"`
	Hints    []string `json:"hints"`
}

// ...existing code...

// loadAnimals loads the animal list from a JSON file.
// Returns a slice of Animal structs or an error if loading/parsing fails.
func loadAnimals(filename string) ([]Animal, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var animals []Animal
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&animals); err != nil {
		return nil, err
	}
	return animals, nil
}

// playGame runs the terminal-based animal guessing game.
func playGame(animals []Animal) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	scanner := bufio.NewScanner(os.Stdin)

	for {
		animal := animals[r.Intn(len(animals))]
		hintIndex := 0
		fmt.Println("\n--- Animal Guessing Game ---")
		fmt.Println("Guess the animal! Type your answer and press Enter.")

		for {
			fmt.Printf("Hint %d: %s\n", hintIndex+1, animal.Hints[hintIndex])
			fmt.Print("Your guess: ")
			scanner.Scan()
			guess := strings.ToLower(strings.TrimSpace(scanner.Text()))

			accepted := map[string]bool{strings.ToLower(animal.Name): true}
			for _, alt := range animal.AltNames {
				accepted[strings.ToLower(alt)] = true
			}

			if accepted[guess] {
				fmt.Printf("Correct! The animal was: %s\n", animal.Name)
				break
			} else {
				hintIndex++
				if hintIndex >= len(animal.Hints) {
					fmt.Printf("Out of hints! The animal was: %s\n", animal.Name)
					break
				} else {
					fmt.Println("Incorrect! Try again.")
				}
			}
		}

		fmt.Print("\nPlay again? (y/n): ")
		scanner.Scan()
		resp := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if resp != "y" && resp != "yes" {
			fmt.Println("Thanks for playing!")
			return
		}
	}
}

// ...all web server and HTTP handler code removed...

func main() {
	// Seed the random number generator for true randomness
	rand.Seed(time.Now().UnixNano())

	// Load animals from the JSON file
	animals, err := loadAnimals("animals.json")
	if err != nil {
		fmt.Println("Error loading animals:", err)
		return
	}
	if len(animals) == 0 {
		fmt.Println("No animals found in animals.json.")
		return
	}

	playGame(animals)
}
