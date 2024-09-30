package Hangman

import (
    "fmt"
)

// Fonction pour lancer la partie de hangman
func Finalworld() {
    fileName := ChooseDifficulty()

    wordList, err := GetRandomWord(fileName)
    if err != nil {
        fmt.Println("Erreur lors du chargement des mots :", err)
        return
    }

    chosenWord := ChooseRandomWord(wordList)

    blanks := Blankword(chosenWord)
    fmt.Println("Mot à deviner : ", blanks)

    PlayGame(chosenWord, blanks)
}