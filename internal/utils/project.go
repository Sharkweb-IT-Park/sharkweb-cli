package utils

import (
	"fmt"
	"os"
)

func ValidateProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dir + "/sharkweb.config.yaml"); os.IsNotExist(err) {
		return "", fmt.Errorf("⛔ Not a Sharkweb project (missing sharkweb.config.yaml)")
	}

	return dir, nil
}
