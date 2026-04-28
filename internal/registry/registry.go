package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Module struct {
	Repo         string   `json:"repo"`
	Version      string   `json:"version"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Dependencies []string `json:"dependencies"`
	Backend      bool     `json:"backend"`
	Frontend     bool     `json:"frontend"`
}

type Registry struct {
	Modules map[string]Module `json:"modules"`
}

var registryURL = "https://raw.githubusercontent.com/Sharkweb-IT-Park/sharkweb-registry/main/registry.json" // Change this to Actual url.

func FetchRegistry() (*Registry, error) {
	resp, err := http.Get(registryURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var reg Registry
	if err := json.NewDecoder(resp.Body).Decode(&reg); err != nil {
		return nil, err
	}

	return &reg, nil
}

func GetModule(name string) (*Module, error) {
	reg, err := FetchRegistry()
	if err != nil {
		return nil, err
	}

	module, ok := reg.Modules[name]
	if !ok {
		return nil, fmt.Errorf("Module not found: %s", name)
	}
	return &module, nil
}
