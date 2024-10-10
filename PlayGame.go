package Hangman

import (
    "fmt"
    "strings"
)

// Fonction pour gérer la boucle de jeu
func PlayGame(chosenWord, blanks string) {
    errors := 0
    chances := 10
    guessedLetters := make(map[string]bool)
    for {
        letter := ProposeLetter(guessedLetters) 
        guessedLetters[letter] = true 
        fmt.Println("Lettre proposée :", letter)

        if !strings.Contains(chosenWord, letter) {
            errors ++
            fmt.Printf("Mauvaise lettre. Il te restes %d tentative(s) ! \n", chances-errors)
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
            fmt.Println("Tu as perdu !")
            break
        }
    }
}

