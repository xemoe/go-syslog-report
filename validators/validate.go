package validators

import (
	"os"
)

func ValidateFileIsExists(file string) (bool, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false, err
	} else {
		return true, nil
	}
}
func ValidateFilesIsExists(files []string) (bool, error) {
	expected := len(files)
	counter := 0
	for _, file := range files {
		if ok, err := ValidateFileIsExists(file); !ok {
			return false, err
		} else {
			counter = counter + 1
		}
	}

	return counter == expected, nil
}
