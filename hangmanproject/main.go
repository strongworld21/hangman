package main

import (
	"fmt"
	"hangman"
	"math/rand"
)

func main() {
	var level string
	fmt.Println("Choisissez un niveau de difficulté (facile, moyen, difficile) :")
	fmt.Scanln(&level)

	var fileName string
	switch level {
	case "facile":
		fileName = "words.txt"
	case "moyen":
		fileName = "words2.txt"
	case "difficile":
		fileName = "words3.txt"
	default:
		fmt.Println("Niveau invalide, veuillez choisir facile, moyen, ou difficile.")
		return
	}
	{
		words, err := hangman.GetRandomWord(fileName)
		if err != nil {
			fmt.Println("Erreur lors du chargement des mots :", err)
			return
		}
		randomIndex := rand.Intn(len(words))
		chosenWord := words[randomIndex]
		fmt.Println(chosenWord)
	}
}
