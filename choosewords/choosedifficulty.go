package Hangman

import "fmt"

// Fonction pour choisir le niveau de difficulté
func ChooseDifficulty() string {
	var FileName string
	var level string
	fmt.Println("Choisissez un niveau de difficulté (facile, moyen, difficile) :")
	fmt.Scanln(&level)

	switch level {
	case "facile":
		FileName = "words.txt"
	case "moyen":
		FileName = "words2.txt"
	case "difficile":
		FileName = "words3.txt"
	default:
		fmt.Println("Niveau invalide, veuillez choisir : facile, moyen, ou difficile.")
	}

	return FileName
}
