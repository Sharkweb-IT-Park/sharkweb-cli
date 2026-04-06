package module

import (
	"fmt"
	"os"
	"path/filepath"
	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/registry"
	"sharkweb-cli/internal/wiring"
)

func InstallModule(name string, projectRoot string, installed map[string]bool) error {

	// =========================
	// 🛑 PREVENT DUPLICATE (RUNTIME)
	// =========================
	if installed[name] {
		return nil
	}

	fmt.Println("🔍 Resolving module:", name)

	// =========================
	// 📦 FETCH MODULE META
	// =========================
	moduleMeta, err := registry.GetModule(name)
	if err != nil {
		return fmt.Errorf("failed to fetch module: %w", err)
	}

	// =========================
	// 🔁 RESOLVE DEPENDENCIES FIRST
	// =========================
	for _, dep := range moduleMeta.Dependencies {
		err := InstallModule(dep, projectRoot, installed)
		if err != nil {
			return err
		}
	}

	fmt.Println("🚀 Installing module:", name)

	// =========================
	// 📥 CLONE MODULE
	// =========================
	tempDir := filepath.Join(os.TempDir(), "sharkweb-"+name)
	_ = os.RemoveAll(tempDir)

	err = CloneModule(moduleMeta.Repo, tempDir)
	if err != nil {
		return fmt.Errorf("clone failed: %w", err)
	}

	// =========================
	// ⚙️ PARSE module.yaml
	// =========================
	moduleConfig, err := ParseModuleConfig(tempDir)
	if err != nil {
		return fmt.Errorf("invalid module.yaml: %w", err)
	}

	// =========================
	// 🔥 VALIDATION
	// =========================
	if moduleConfig.Backend && moduleConfig.Entry.Backend == "" {
		return fmt.Errorf("backend entry missing in module.yaml")
	}

	if moduleConfig.Frontend && moduleConfig.Entry.Frontend == "" {
		return fmt.Errorf("frontend entry missing in module.yaml")
	}

	fmt.Println("📦 Module:", moduleConfig.Name)

	// =========================
	// 📦 BACKEND INSTALL
	// =========================
	if moduleConfig.Backend {
		src := filepath.Join(tempDir, "backend")
		dst := filepath.Join(projectRoot, "backend/modules", name)

		if exists(src) {
			err = CopyDir(src, dst)
			if err != nil {
				return err
			}
			fmt.Println("✅ Backend installed")
		}
	}

	// =========================
	// 🎨 FRONTEND INSTALL
	// =========================
	if moduleConfig.Frontend {
		src := filepath.Join(tempDir, "frontend")
		dst := filepath.Join(projectRoot, "frontend/modules", name)

		if exists(src) {
			err = CopyDir(src, dst)
			if err != nil {
				return err
			}
			fmt.Println("✅ Frontend installed")
		}
	}

	// =========================
	// 🔗 SHARED LAYER MERGE
	// =========================
	sharedSrc := filepath.Join(tempDir, "shared")
	sharedDst := filepath.Join(projectRoot, "shared")

	if exists(sharedSrc) {
		err = CopyDirSafe(sharedSrc, sharedDst)
		if err != nil {
			return err
		}
		fmt.Println("🔗 Shared layer merged")
	}

	// =========================
	// ⚙️ LOAD CONFIG
	// =========================
	cfg, err := config.Load(projectRoot)
	if err != nil {
		return err
	}

	// =========================
	// 🛑 PREVENT DUPLICATE (CONFIG LEVEL)
	// =========================
	if config.IsModuleInstalled(cfg, name) {
		fmt.Println("⚠️ Module already in config:", name)
	} else {
		// =========================
		// ➕ ADD MODULE
		// =========================
		config.AddModule(cfg, name)

		// =========================
		// 💾 SAVE CONFIG
		// =========================
		err = config.Save(projectRoot, cfg)
		if err != nil {
			return err
		}
	}

	// =========================
	// 🔥 AUTO-WIRING
	// =========================
	err = wiring.GenerateWiring(projectRoot, cfg.Modules)
	if err != nil {
		return err
	}

	// =========================
	// ✅ MARK INSTALLED
	// =========================
	installed[name] = true

	fmt.Println("✅ Installed:", name)

	return nil
}

// =========================
// 📁 HELPERS
// =========================

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CopyDirSafe(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, os.ModePerm)
		}

		// skip if exists
		if _, err := os.Stat(target); err == nil {
			return nil
		}

		return CopyFile(path, target)
	})
}
