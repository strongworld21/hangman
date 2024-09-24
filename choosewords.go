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