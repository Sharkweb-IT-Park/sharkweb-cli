package module

import (
	"fmt"
	"os"
	"path/filepath"

	"sharkweb-cli/internal/registry"
)

// =========================
// 🔹 INSTALL CORE (RECURSIVE)
// =========================
func InstallModule(
	name string,
	projectRoot string,
	visited map[string]bool,
	installed map[string]bool,
) error {

	// 🔁 prevent infinite recursion
	if visited[name] {
		return nil
	}
	visited[name] = true

	fmt.Println("🔍 Resolving module:", name)

	// =========================
	// 📦 FETCH META
	// =========================
	moduleMeta, err := registry.GetModule(name)
	if err != nil {
		return fmt.Errorf("failed to fetch module: %w", err)
	}

	if moduleMeta.Repo == "" {
		return fmt.Errorf("invalid repo for module: %s", name)
	}

	// =========================
	// 🔁 INSTALL DEPENDENCIES FIRST
	// =========================
	for _, dep := range moduleMeta.Dependencies {
		if err := InstallModule(dep, projectRoot, visited, installed); err != nil {
			return err
		}
	}

	fmt.Println("🚀 Installing module:", name)

	// =========================
	// 📥 CLONE MODULE
	// =========================
	tempDir := filepath.Join(os.TempDir(), "sharkweb-"+name)
	_ = os.RemoveAll(tempDir)

	if err := CloneModule(moduleMeta.Repo, tempDir); err != nil {
		return fmt.Errorf("clone failed: %w", err)
	}

	// =========================
	// ⚙️ PARSE CONFIG
	// =========================
	moduleConfig, err := ParseModuleConfig(tempDir)
	if err != nil {
		return fmt.Errorf("invalid module.yaml: %w", err)
	}

	// =========================
	// 🔥 VALIDATION
	// =========================
	if moduleConfig.Backend && moduleConfig.Entry.Backend == "" {
		return fmt.Errorf("backend entry missing")
	}
	if moduleConfig.Frontend && moduleConfig.Entry.Frontend == "" {
		return fmt.Errorf("frontend entry missing")
	}

	// =========================
	// 📦 BACKEND COPY
	// =========================
	if moduleConfig.Backend {
		src := filepath.Join(tempDir, "backend")
		dst := filepath.Join(projectRoot, "backend/modules", name)

		if exists(dst) {
			return fmt.Errorf("module already exists: %s", name)
		}

		if !exists(src) {
			return fmt.Errorf("backend folder missing in module: %s", name)
		}

		if err := CopyDir(src, dst); err != nil {
			return err
		}

		fmt.Println("✅ Backend installed")
	}

	// =========================
	// 🎨 FRONTEND COPY
	// =========================
	if moduleConfig.Frontend {
		src := filepath.Join(tempDir, "frontend")
		dst := filepath.Join(projectRoot, "frontend/modules", name)

		if exists(dst) {
			return fmt.Errorf("module already exists: %s", name)
		}

		if !exists(src) {
			return fmt.Errorf("frontend folder missing in module: %s", name)
		}

		if err := CopyDir(src, dst); err != nil {
			return err
		}

		fmt.Println("✅ Frontend installed")
	}

	// =========================
	// 🔗 SHARED COPY
	// =========================
	sharedSrc := filepath.Join(tempDir, "shared")
	sharedDst := filepath.Join(projectRoot, "shared")

	if exists(sharedSrc) {
		if err := CopyDirSafe(sharedSrc, sharedDst); err != nil {
			return err
		}
		fmt.Println("🔗 Shared merged")
	}

	// ✅ MARK SUCCESSFUL INSTALL
	installed[name] = true

	fmt.Println("✅ Installed:", name)

	return nil
}
