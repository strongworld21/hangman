package Hangman

import (
    "fmt"
)

// Fonction pour proposer une lettre et la lire dans le terminal
func ProposeLetter(guessedLetters map[string]bool) string {
    var letter string
    for {
        fmt.Print("Propose une lettre : ")
        fmt.Scan(&letter)
        if len(letter) == 1 && !guessedLetters[letter] {
            return letter
        } else if guessedLetters[letter] {
            fmt.Println("Vous avez déjà proposé cette lettre. Essayez une autre lettre.")
        } else {
            fmt.Println("Veuillez proposer une seule lettre.")
        }
    }
}
