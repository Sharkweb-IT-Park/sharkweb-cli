package module

import (
	"fmt"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/wiring"
)

func AddModule(name string, projectRoot string) error {

	visited := map[string]bool{}
	installed := map[string]bool{}

	// =========================
	// 📦 INSTALL MODULES
	// =========================
	if err := InstallModule(name, projectRoot, visited, installed); err != nil {
		return err
	}

	// =========================
	// 📥 LOAD CONFIG
	// =========================
	cfg, err := config.Load(projectRoot)
	if err != nil {
		return err
	}

	// =========================
	// ➕ ADD INSTALLED MODULES
	// =========================
	for m := range installed {
		if !config.IsModuleInstalled(cfg, m) {
			config.AddModule(cfg, m)
		}
	}

	fmt.Println("📦 Final modules:", cfg.Modules)

	// =========================
	// 💾 SAVE CONFIG
	// =========================
	if err := config.Save(projectRoot, cfg); err != nil {
		return err
	}

	// =========================
	// 🔧 GENERATE WIRING
	// =========================
	if err := wiring.GenerateBackendWiring(projectRoot, cfg.Modules); err != nil {
		return err
	}

	if err := wiring.GenerateFrontendWiring(projectRoot, cfg.Modules); err != nil {
		return err
	}

	if err := wiring.GenerateNextRoutes(projectRoot, cfg.Modules); err != nil {
		return err
	}

	fmt.Println("🚀 Module installed & wired successfully")

	return nil
}
