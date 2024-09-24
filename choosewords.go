package hangman

import (
    "os"
	"strings"
    "fmt"
)

func GetRandomWord(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(data), "\n")
    var words []string

    for _, line := range lines {
        trimmedLine := strings.TrimSpace(line)
        if trimmedLine != "" {
            words = append(words, trimmedLine)
        }
    }

    return words, nil
}
func AskForLetter() {
    fmt.Printf("words %s letter : ", strings.Join(blanks, ""))
    var input string
    fmt.Println(&input)
    fmt.Println(input)
}