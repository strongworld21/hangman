package main

import (
	"Hangman/choosewords"
	"fmt"
)

func main() {
	// Choisir la difficulté
	fileName := Hangman.ChooseDifficulty()

	// Lire les mots du fichier avec gestion d'erreur
	wordList, err := Hangman.GetRandomWord(fileName)
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	Hangman.ChooseRandomWord(wordList)
}

