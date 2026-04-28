package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name    string   `yaml:"name"`
	Modules []string `yaml:"modules"`
}

// =========================
// 📥 LOAD CONFIG
// =========================
func Load(projectRoot string) (*Config, error) {

	path := filepath.Join(projectRoot, "sharkweb.config.yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		// ✅ IMPORTANT: DO NOT FAIL
		return &Config{
			Name:    "sharkweb-app",
			Modules: []string{},
		}, nil
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Modules == nil {
		cfg.Modules = []string{}
	}

	return &cfg, nil
}

// =========================
// 💾 SAVE CONFIG
// =========================
func Save(projectRoot string, cfg *Config) error {

	path := filepath.Join(projectRoot, "sharkweb.config.yaml") // ✅ HERE

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// =========================
// 🔍 CHECK MODULE
// =========================
func IsModuleInstalled(cfg *Config, name string) bool {
	for _, m := range cfg.Modules {
		if m == name {
			return true
		}
	}
	return false
}

// =========================
// ➕ ADD MODULE
// =========================
func AddModule(cfg *Config, name string) {
	if !IsModuleInstalled(cfg, name) {
		cfg.Modules = append(cfg.Modules, name)
	}
}

// =========================
// ➖ REMOVE MODULE
// =========================

func RemoveModule(cfg *Config, name string) {
	var updated []string

	for _, m := range cfg.Modules {
		if m != name {
			updated = append(updated, m)
		}
	}

	cfg.Modules = updated
}
