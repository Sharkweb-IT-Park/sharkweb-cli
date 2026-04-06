package project

import (
	"os"

	"gopkg.in/yaml.v3"
)

func CreateConfig(projectRoot string) error {

	cfg := map[string]interface{}{
		"name":    "sharkweb-app",
		"modules": []string{},
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(projectRoot+"/sharkweb.config.yaml", data, 0644)
}
