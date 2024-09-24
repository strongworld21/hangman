package main

import (
	"Hangman/choosewords"
	"fmt"
	"math/rand"
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

	// Choisir un mot aléatoire
	randomIndex := rand.Intn(len(wordList))
	chosenWord := wordList[randomIndex]
	fmt.Println("Mot choisi :", chosenWord)
}

