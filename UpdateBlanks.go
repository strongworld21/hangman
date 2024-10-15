package Hangman

import (
	"strings"
)

// Fonction pour mettre à jour les espaces vides avec les lettres trouvées
func UpdateBlanks(chosenWord, blanks, letter string) string {
	updatedBlanks := strings.Fields(blanks)

	for i, ch := range chosenWord {
		if string(ch) == letter {
			updatedBlanks[i] = string(ch)
		}
	}
	return strings.Join(updatedBlanks, " ")
}
