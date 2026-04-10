package wiring

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func toPascalCase(s string) string {
	parts := strings.Split(s, "-")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func toCamelCase(s string) string {
	p := toPascalCase(s)
	return strings.ToLower(p[:1]) + p[1:]
}

func indent(s string, spaces int) string {
	prefix := strings.Repeat(" ", spaces)
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		lines[i] = prefix + line
	}

	return strings.Join(lines, "\n")
}

func GetBackendModulePath(projectRoot string) (string, error) {

	goModPath := filepath.Join(projectRoot, "backend", "go.mod")

	file, err := os.Open(goModPath)
	if err != nil {
		return "", fmt.Errorf("cannot open go.mod at %s: %w", goModPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module not found in %s", goModPath)
}
