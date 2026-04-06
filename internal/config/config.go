package config

import (
	"os"

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
	path := projectRoot + "/sharkweb.config.yaml"

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// =========================
// 💾 SAVE CONFIG
// =========================
func Save(projectRoot string, cfg *Config) error {
	path := projectRoot + "/sharkweb.config.yaml"

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
