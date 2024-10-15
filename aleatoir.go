package Hangman

import (
	"math/rand"
)

func ChooseRandomWord(wordList []string) string {
	randomIndex := rand.Intn(len(wordList))
	chosenWord := wordList[randomIndex]
	return chosenWord
}
