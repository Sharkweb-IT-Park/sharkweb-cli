package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ToPascalCase(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func getTemplatePath() (string, error) {

	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// Go up from binary location
	baseDir := filepath.Dir(execPath)

	return filepath.Join(baseDir, "internal", "generator", "templates", "module", "full"), nil
}

func GetModules(projectRoot string) ([]string, error) {

	modulesPath := filepath.Join(projectRoot, "backend/modules")

	entries, err := os.ReadDir(modulesPath)
	if err != nil {
		return nil, err
	}

	var modules []string

	for _, entry := range entries {
		if entry.IsDir() {
			modules = append(modules, entry.Name())
		}
	}

	return modules, nil
}

func GetGoModuleName(projectRoot string) (string, error) {

	goModPath := filepath.Join(projectRoot, "go.mod")

	file, err := os.Open(goModPath)
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "module ") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module "))
			return moduleName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module name not found in go.mod")
}
