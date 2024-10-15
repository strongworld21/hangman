package Hangman

import (
	"fmt"
)

func ChooseDifficulty() (string, string) {
	var fileName string
	var level string
	fmt.Println("Choisissez un niveau de difficulté (facile, moyen, difficile) :")
	fmt.Scanln(&level)

	switch level {
	case "facile":
		fileName = "words.txt"
	case "moyen":
		fileName = "words2.txt"
	case "difficile":
		fileName = "words3.txt"
	default:
		fmt.Println("Niveau invalide, veuillez choisir : facile, moyen, ou difficile.")
	}

	return fileName, level
}
