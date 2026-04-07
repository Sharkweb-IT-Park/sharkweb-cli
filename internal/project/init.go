package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitProject(name string) (string, error) {

	root := filepath.Join(".", name)

	err := os.MkdirAll(root, os.ModePerm)
	if err != nil {
		return "", err
	}

	// create folders
	// os.MkdirAll(filepath.Join(root, "backend"), os.ModePerm)
	// os.MkdirAll(filepath.Join(root, "frontend"), os.ModePerm)
	// os.MkdirAll(filepath.Join(root, "shared"), os.ModePerm)

	fmt.Println("📁 Project structure created")

	return root, nil
}
