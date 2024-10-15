package Hangman

import (
	"fmt"
	"strings"
)

func PlayGame(chosenWord, blanks string) {
	errors := 0
	chances := 10
	guessedLetters := make(map[string]bool)
	for {
		letter := ProposeLetter(guessedLetters)
		guessedLetters[letter] = true
		fmt.Println("Lettre proposée :", letter)

		if !strings.Contains(chosenWord, letter) {
			errors++
			PrintHangman(errors)
			fmt.Printf("Mauvaise lettre. Il te reste %d tentative(s) ! \n", chances-errors)
			fmt.Println("Mot actuel : ", blanks)
		} else {
			blanks = UpdateBlanks(chosenWord, blanks, letter)
			fmt.Println("Mot actuel : ", blanks)
			fmt.Println("Bonne lettre !")
		}

		if strings.ReplaceAll(blanks, " ", "") == chosenWord {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", chosenWord)
			break
		}
		if errors == chances {
			fmt.Println("Tu as perdu ! Le mot était :", chosenWord)
			PrintHangman(errors)
			break
		}
	}
}
