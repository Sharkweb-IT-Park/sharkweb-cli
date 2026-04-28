package publish

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExtractModule(moduleName string) (string, error) {

	root, err := os.Getwd()
	if err != nil {
		return "", err
	}

	tmpDir := filepath.Join(os.TempDir(), "sharkweb-"+moduleName)

	// 🔥 clean old temp
	_ = os.RemoveAll(tmpDir)

	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return "", err
	}

	// =========================
	// 🔹 SOURCE PATHS
	// =========================
	backendSrc := filepath.Join(root, "backend/modules", moduleName)
	frontendSrc := filepath.Join(root, "frontend/modules", moduleName)
	sharedSrc := filepath.Join(root, "shared/modules", moduleName)

	// =========================
	// 🔹 COPY MODULE PARTS
	// =========================
	copied := false

	if exists(backendSrc) {
		if err := copyDir(backendSrc, filepath.Join(tmpDir, "backend")); err != nil {
			return "", err
		}
		copied = true
	}

	if exists(frontendSrc) {
		if err := copyDir(frontendSrc, filepath.Join(tmpDir, "frontend")); err != nil {
			return "", err
		}
		copied = true
	}

	if exists(sharedSrc) {
		if err := copyDir(sharedSrc, filepath.Join(tmpDir, "shared")); err != nil {
			return "", err
		}
		copied = true
	}

	if !copied {
		return "", fmt.Errorf("module '%s' not found in backend/frontend/shared", moduleName)
	}

	// =========================
	// 🔹 FIND module.yaml (SMART)
	// =========================

	possiblePaths := []string{
		filepath.Join(backendSrc, "module.yaml"),                  // ✅ recommended
		filepath.Join(root, "modules", moduleName, "module.yaml"), // optional central
		filepath.Join(frontendSrc, "module.yaml"),                 // fallback
		filepath.Join(sharedSrc, "module.yaml"),                   // fallback
	}

	var foundYaml string

	for _, p := range possiblePaths {
		if exists(p) {
			foundYaml = p
			break
		}
	}

	if foundYaml == "" {
		return "", fmt.Errorf("module.yaml not found for module '%s'", moduleName)
	}

	// copy module.yaml
	if err := copyFile(foundYaml, filepath.Join(tmpDir, "module.yaml")); err != nil {
		return "", err
	}

	return tmpDir, nil
}
