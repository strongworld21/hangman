package Hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fonction pour proposer une lettre et la lire dans le terminal
func ProposeLetter() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Propose une lettre : ")
		letter, _ := reader.ReadString('\n')     
		letter = strings.TrimSpace(letter)        
		if len(letter) == 1 && letter >= "a" && letter <= "z" || letter >= "A" && letter <= "Z" {
			return letter 
		}
		fmt.Println("Merci de proposer une seule lettre alphabétique.")
	}
}
