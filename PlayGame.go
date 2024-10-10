package Hangman

import (
    "fmt"
    "strings"
)

// Fonction pour gérer la boucle de jeu
func PlayGame(chosenWord, blanks string) {
    guessedLetters := make(map[string]bool)
    for {
        letter := ProposeLetter(guessedLetters) 
        guessedLetters[letter] = true 
        fmt.Println("Lettre proposée :", letter)

        if strings.Contains(chosenWord, letter) {
            fmt.Println("Bonne lettre !")
            blanks = UpdateBlanks(chosenWord, blanks, letter) 
            fmt.Println("Mot actuel : ", blanks)
        } else {
            fmt.Println("Mauvaise lettre.")
        }

        if strings.ReplaceAll(blanks, " ", "") == chosenWord {
            fmt.Println("Félicitations ! Vous avez deviné le mot :", chosenWord)
            break
        }
    }
}
