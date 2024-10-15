package Hangman

import (
	"strings"
)

func BlankWord(chosenWord string) string {
	blanks := []string{}
	for range chosenWord {
		blanks = append(blanks, "_")
	}
	return (strings.Join(blanks, " "))
}
