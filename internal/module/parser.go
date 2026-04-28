package module

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ModuleConfig struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Backend      bool     `yaml:"backend"`
	Frontend     bool     `yaml:"frontend"`
	Dependencies []string `yaml:"dependencies"`

	Entry struct {
		Backend  string `yaml:"backend"`
		Frontend string `yaml:"frontend"`
	} `yaml:"entry"`
}

// =========================
// 🔹 PARSE module.yaml
// =========================
func ParseModuleConfig(dir string) (*ModuleConfig, error) {

	path := filepath.Join(dir, "module.yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("module.yaml not found")
	}

	var cfg ModuleConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Name == "" {
		return nil, fmt.Errorf("invalid module.yaml (missing name)")
	}

	return &cfg, nil
}
