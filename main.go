package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var difficulty uint8
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println(`
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)
	fmt.Scanln(&difficulty)

	var tries uint8
	switch difficulty {
	case 1:
		tries = 10
	case 2:
		tries = 5
	case 3:
		tries = 3
	default:
		return
	}

	var search, guess uint8
	search = uint8(rand.Intn(100))

	fmt.Println("Enter your choice: ")
	fmt.Scanln(&guess)

	for {
		tries--

		if guess == search {
			fmt.Println("You WOn")
			return
		}

		if tries > 0 {
			fmt.Println("Wrong number. Try again: ")
			fmt.Scanln(&guess)
		} else {
			break
		}

	}

	fmt.Println("You lost loser")
	fmt.Scanln()
}
