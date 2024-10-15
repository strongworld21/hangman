package main

import (
	"Hangman"
	"fmt"
)

func main() {
	fileName, level := Hangman.ChooseDifficulty()

	wordList, err := Hangman.GetRandomWord(fileName)
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	chosenWord := Hangman.ChooseRandomWord(wordList)

	var blanks string
	if level == "facile" {
		blanks = Hangman.LettersRevealed(chosenWord)
	} else {
		blanks = Hangman.BlankWord(chosenWord)
	}

	fmt.Println("Mot à deviner : ", blanks)

	Hangman.PlayGame(chosenWord, blanks)
}
