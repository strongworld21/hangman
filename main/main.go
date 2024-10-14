package main

import (
	"Hangman"
	"fmt"
)

// Fonction pour lancer la partie de hangman
func main() {
    fileName := Hangman.ChooseDifficulty()

    wordList, err := Hangman.GetRandomWord(fileName)
    if err != nil {
        fmt.Println("Erreur lors du chargement des mots :", err)
        return
    }

    chosenWord := Hangman.ChooseRandomWord(wordList)

    blanks := Hangman.Blankword(chosenWord)
    fmt.Println("Mot à deviner : ", blanks)

    Hangman.PlayGame(chosenWord, blanks)
}
