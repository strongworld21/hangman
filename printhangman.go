package Hangman

import "fmt"


func PrintHangman(errors int) {
    hangman := []string{
        ``,
		    `



        


            =========`,	
        `

              |
              |
              |
              |
              |
            =========`,
        `
          +---+
              |
              |
              |
              |
              |
            =========`,
        `
          +---+
          |   |
              |
              |
              |
              |
            =========`,
        `
          +---+
          |   |
          O   |
              |
              |
              |
            =========`,
        `
          +---+
          |   |
          O   |
          |   |
              |
              |
            =========`,
        `
          +---+
          |   |
          O   |
         /|   |
              |
              |
            =========`,
        `
          +---+
          |   |
          O   |
         /|\  |
              |
              |
            =========`,
        `
          +---+
          |   |
          O   |
         /|\  |
         /    |
              |
            =========`,
        `
          +---+
          |   |
          O   |
         /|\  |
         / \  |
              |
            =========`,
    }

    if errors >= 0 && errors < len(hangman) {
        fmt.Println(hangman[errors])
    }
}
