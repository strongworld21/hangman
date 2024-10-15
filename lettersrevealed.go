package Hangman

import (
	"math/rand"
	"strings"
)

func LettersRevealed(chosenWord string) string {
	length := len(chosenWord)
	lettersToReveal := (length / 2) - 1
	if lettersToReveal <= 0 { //si mot de 2 lettres
		lettersToReveal = 1
	}

	revealedIndices := make(map[int]bool)

	for len(revealedIndices) < lettersToReveal {
		randomIndex := rand.Intn(length)
		revealedIndices[randomIndex] = true
	}
	blanks := []string{}
	for i, ch := range chosenWord {
		if revealedIndices[i] {
			blanks = append(blanks, string(ch))
		} else {
			blanks = append(blanks, "_")
		}
	}

	return strings.Join(blanks, " ")
}
