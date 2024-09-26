package Hangman

import (
	"fmt"
)

func Finalworld() {
	// Choisir la difficulté
	fileName := ChooseDifficulty()

	// Lire les mots du fichier avec gestion d'erreur
	wordList, err := GetRandomWord(fileName)
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	chosenWord := ChooseRandomWord(wordList)
	blanks := Blankword(chosenWord)
	fmt.Println(blanks)
}
