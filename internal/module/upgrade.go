package module

import (
	"fmt"

	"sharkweb-cli/internal/config"
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
	err = RemoveModule(name, projectRoot)
	if err != nil {
		return err
	}

	// =========================
	// 🚀 INSTALL LATEST VERSION
	// =========================
	installed := make(map[string]bool)

	err = InstallModule(name, projectRoot, installed)
	if err != nil {
		return err
	}

	fmt.Println("✅ Module upgraded:", name)

	return nil
}
