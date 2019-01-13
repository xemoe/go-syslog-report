package parser

import (
	"os"
)

func ValidateFilesIsExists(files []string) (bool, error) {
	expected := len(files)
	counter := 0
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return false, err
		} else {
			counter = counter + 1
		}
	}

	return counter == expected, nil
}
