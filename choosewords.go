package hangman

import (
    "os"
	"strings"
)

func GetRandomWord(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
    if err != nil {
        return nil, err
    }
	words := strings.Split(string(data), "\n")

    return words, nil
}
