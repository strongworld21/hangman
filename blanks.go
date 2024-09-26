package Hangman

import (
	"strings"
)

func Blankword(chosenWord string) string {
	blanks := []string{}
	for range chosenWord {
		blanks = append(blanks, "_")
	}
	return (strings.Join(blanks, " "))
}
