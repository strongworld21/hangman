package Hangman

import (
	"math/rand"
)

// Fonction pour choisir un mot aléatoire à partir d'une liste
func ChooseRandomWord(wordList []string) string {
	randomIndex := rand.Intn(len(wordList))
	chosenWord := wordList[randomIndex]
	return chosenWord
}
