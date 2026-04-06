package module

import (
	"fmt"
	"os"
	"path/filepath"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/wiring"
)

func RemoveModule(name string, projectRoot string) error {

	fmt.Println("🗑 Removing module:", name)

	// =========================
	// 📁 REMOVE BACKEND
	// =========================
	backendPath := filepath.Join(projectRoot, "backend/modules", name)

	if exists(backendPath) {
		err := os.RemoveAll(backendPath)
		if err != nil {
			return err
		}
		fmt.Println("✅ Backend removed")
	}

	// =========================
	// 🎨 REMOVE FRONTEND
	// =========================
	frontendPath := filepath.Join(projectRoot, "frontend/modules", name)

	if exists(frontendPath) {
		err := os.RemoveAll(frontendPath)
		if err != nil {
			return err
		}
		fmt.Println("✅ Frontend removed")
	}

	// =========================
	// ⚙️ UPDATE CONFIG
	// =========================
	cfg, err := config.Load(projectRoot)
	if err != nil {
		return err
	}

	config.RemoveModule(cfg, name)

	err = config.Save(projectRoot, cfg)
	if err != nil {
		return err
	}

	// =========================
	// 🔥 REGENERATE WIRING
	// =========================
	err = wiring.GenerateWiring(projectRoot, cfg.Modules)
	if err != nil {
		return err
	}

	fmt.Println("✅ Module removed:", name)

	return nil
}
