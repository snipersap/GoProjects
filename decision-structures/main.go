package main

import (
	"log"
	"strings"
)

// Experiments with decision structures
// if-else, switch-case

func main() {

	animals := []string{"cat", "dog", "elephant", "giraffe", "Whale"}

	for _, animal := range animals {
		switch animal {
		case "cat":
			log.Println("It's a cat!")
		case "dog":
			log.Println("It's a dog!")
		case "elephant":
			log.Println("It's an elephant!")
		default:
			log.Printf("Is %s an animal?\n", animal)
		}
	}

	if animals[4] == "whale" { //If is case sensitive
		log.Println("It's a whale!")
	} else {
		log.Println("No idea what animal is this!")
	}

	// Case insensitive comparison. EqualFold uses case folding see wikipedia:
	// https://en.wikipedia.org/wiki/Letter_case#Unicode_case_folding_and_script_identification
	if strings.EqualFold(animals[4], "whale") {
		log.Println("It's a whale!")
	}

}
