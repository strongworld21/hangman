package Hangman

import (
	"fmt"
	"math/rand"
)

// Fonction pour choisir un mot aléatoire à partir d'une liste
func ChooseRandomWord(wordList []string) {
	randomIndex := rand.Intn(len(wordList))
	chosenWord := wordList[randomIndex]
	fmt.Println("Mot choisi :", chosenWord)
}