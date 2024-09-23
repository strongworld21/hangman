package main

import (
    "fmt"
    "hangman"
    "math/rand"
)

func main() {
    words, err := hangman.GetRandomWord("words.txt")
    if err != nil {
        fmt.Println("Erreur lors du chargement des mots :", err)
        return
    }
    randomIndex := rand.Intn(len(words)) 
    chosenWord := words[randomIndex]
    fmt.Println(chosenWord)
}