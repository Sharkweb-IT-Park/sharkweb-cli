package module

import (
	"fmt"

	"sharkweb-cli/internal/registry"
)

func ResolverDependencies(name string, installed map[string]bool) error {
	if installed[name] {
		return nil
	}

	module, err := registry.GetModule(name)
	if err != nil {
		return err
	}

	for _, dep := range module.Dependencies {
		err := ResolverDependencies(dep, installed)
		if err != nil {
			return err
		}
	}

	fmt.Println("📦 Ready to install:", name)
	installed[name] = true
	return nil
}
