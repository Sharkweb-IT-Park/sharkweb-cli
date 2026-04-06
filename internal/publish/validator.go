package publish

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ModuleMeta struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Dependencies []string `yaml:"dependencies"`
	Backend      bool     `yaml:"backend"`
	Frontend     bool     `yaml:"frontend"`
}

func ValidateModule(modulePath string) (*ModuleMeta, error) {

	configPath := filepath.Join(modulePath, "module.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("module.yaml not found")
	}

	var meta ModuleMeta
	err = yaml.Unmarshal(data, &meta)
	if err != nil {
		return nil, err
	}

	if meta.Name == "" || meta.Version == "" {
		return nil, fmt.Errorf("invalid module.yaml")
	}

	return &meta, nil
}
