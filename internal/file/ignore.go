package file

import (
	"bufio"
	"os"
	"strings"
)

func GitignoreContainsKeyEntry() (bool, error) {
	f, err := os.Open(".gitignore")
	if err != nil {
		return false, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "*.key" {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func AppendToGitignore(gitignoreEntry string) error {
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(gitignoreEntry)
	return err
}
