package module

import (
	"fmt"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/wiring"
)

func UpgradeModule(name string, projectRoot string) error {

	fmt.Println("⬆️ Upgrading module:", name)

	// =========================
	// 🔹 LOAD CONFIG
	// =========================
	cfg, err := config.Load(projectRoot)
	if err != nil {
		return err
	}

	// =========================
	// 🔹 CHECK EXISTS
	// =========================
	if !config.IsModuleInstalled(cfg, name) {
		return fmt.Errorf("module not installed: %s", name)
	}

	// =========================
	// 🗑 REMOVE OLD VERSION
	// =========================
	if err := RemoveModule(name, projectRoot); err != nil {
		return err
	}

	// =========================
	// 🚀 INSTALL LATEST VERSION
	// =========================
	visited := map[string]bool{}
	installed := map[string]bool{}

	if err := InstallModule(name, projectRoot, visited, installed); err != nil {
		return err
	}

	// =========================
	// 🔄 RELOAD CONFIG
	// =========================
	cfg, err = config.Load(projectRoot)
	if err != nil {
		return err
	}

	// =========================
	// ➕ ADD MODULE BACK
	// =========================
	for m := range installed {
		if !config.IsModuleInstalled(cfg, m) {
			config.AddModule(cfg, m)
		}
	}

	// =========================
	// 💾 SAVE CONFIG
	// =========================
	if err := config.Save(projectRoot, cfg); err != nil {
		return err
	}

	// =========================
	// 🔧 REGENERATE WIRING
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

	fmt.Println("✅ Module upgraded:", name)

	return nil
}
