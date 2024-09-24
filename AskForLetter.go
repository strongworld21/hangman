package hangman

import (
    "fmt"
	"strings"
)

func AskForLetter() {
    fmt.Printf("words %s letter : ", strings.Join(blanks, ""))
    var input string
    fmt.Println(&input)
    fmt.Println(input)
}