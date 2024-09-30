package Hangman

import (
    "fmt"
    "strings"
)

// Fonction pour gérer la boucle de jeu
func PlayGame(chosenWord, blanks string) {
    guessedLetters := make(map[string]bool)

    // Boucle pour proposer des lettres
    for {
        // Proposer une lettre
        letter := ProposeLetter(guessedLetters) // Assurez-vous que cette fonction est correcte
        guessedLetters[letter] = true // Marque la lettre comme devinée
        fmt.Println("Lettre proposée :", letter)

        // Vérifier si la lettre proposée est dans le mot choisi
        if strings.Contains(chosenWord, letter) {
            fmt.Println("Bonne lettre !")
            // Mettre à jour les blancs avec la lettre trouvée
            blanks = UpdateBlanks(chosenWord, blanks, letter) // Appel à la fonction UpdateBlanks
            fmt.Println("Mot actuel : ", blanks)
        } else {
            fmt.Println("Mauvaise lettre.")
        }

        // Condition d'arrêt si toutes les lettres sont trouvées
        if strings.Join(strings.Fields(blanks), "") == chosenWord {
            fmt.Println("Félicitations ! Vous avez deviné le mot :", chosenWord)
            break
        }
    }
}
