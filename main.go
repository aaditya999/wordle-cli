package main

import (
	"fmt"
	"strings"
)

const (
	targetWord = "STARS"
	maxTries   = 6
	maxLength  = 5
	cols       = "🔳🔳🔳🔳🔳 ."
)

func validateGuess(guess string) string {
	feedback := make([]rune, maxLength*2)
	for i := 0; i < maxLength; i++ {
		if targetWord[i] == guess[i] {
			feedback[i] = '🟩'
			feedback[i+1] = '.'
		} else if strings.ContainsRune(targetWord, rune(guess[i])) {
			feedback[i] = '🟨'
			feedback[i+1] = '.'
		} else {
			feedback[i] = '⬜'
			feedback[i+1] = '.'
		}
	}
	return string(feedback)
}
func main() {
	i := 0
	var guess string
	for j := 1; j <= maxLength; j++ {
		fmt.Println(cols + " ")

	}

	for i <= maxTries {
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println(guess)
		}
		if len(guess) != maxLength {
			fmt.Println("The word should be of 5 letters")
			continue
		}
		fmt.Println(validateGuess(guess))
		for j := i + 2; j <= 6; j++ {
			fmt.Println(cols)
		}
		if guess == targetWord {
			fmt.Println("🎉")
			return
		}
		i++
	}

}
