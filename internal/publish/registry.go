package publish

import (
	"encoding/json"
	"os"
)

type Registry map[string]ModuleMeta

var registryPath = "./registry.json"

func LoadRegistry() (Registry, error) {

	data, err := os.ReadFile(registryPath)
	if err != nil {
		return Registry{}, nil
	}

	var reg Registry
	_ = json.Unmarshal(data, &reg)

	return reg, nil
}

func SaveRegistry(reg Registry) error {

	data, err := json.MarshalIndent(reg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(registryPath, data, 0644)
}

func UpdateRegistry(meta *ModuleMeta, repo string) error {

	reg, _ := LoadRegistry()

	reg[meta.Name] = ModuleMeta{
		Name:         meta.Name,
		Version:      meta.Version,
		Dependencies: meta.Dependencies,
		Backend:      meta.Backend,
		Frontend:     meta.Frontend,
	}

	return SaveRegistry(reg)
}
