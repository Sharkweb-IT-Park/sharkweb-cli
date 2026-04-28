package publish

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ModuleMeta struct {
	Name         string   `yaml:"name" json:"name"`
	Version      string   `yaml:"version" json:"version"`
	Dependencies []string `yaml:"dependencies" json:"dependencies"`
	Backend      bool     `yaml:"backend" json:"backend"`
	Frontend     bool     `yaml:"frontend" json:"frontend"`
	Repo         string   `json:"repo"`
}

func ValidateModule(modulePath string) (*ModuleMeta, error) {

	configPath := filepath.Join(modulePath, "module.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("module.yaml not found")
	}

	var meta ModuleMeta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	if meta.Name == "" {
		return nil, fmt.Errorf("module name required")
	}

	if meta.Version == "" {
		return nil, fmt.Errorf("version required")
	}

	if !meta.Backend && !meta.Frontend {
		return nil, fmt.Errorf("module must have backend or frontend")
	}

	return &meta, nil
}
