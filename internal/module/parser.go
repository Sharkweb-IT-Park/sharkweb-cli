package module

import (
	"os"

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

	Events struct {
		Publish   []string `yaml:"publish"`
		Subscribe []string `yaml:"subscribe"`
	} `yaml:"events"`
}

func ParseModuleConfig(path string) (*ModuleConfig, error) {
	data, err := os.ReadFile(path + "/module.yaml")
	if err != nil {
		return nil, err
	}

	var config ModuleConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
