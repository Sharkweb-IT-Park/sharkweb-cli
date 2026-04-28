package publish

import "fmt"

func PublishModule(moduleName string, repo string, version string) error {

	// 🔹 Extract from monorepo
	tmpDir, err := ExtractModule(moduleName)
	if err != nil {
		return fmt.Errorf("extract failed: %w", err)
	}

	// 🔹 Validate module.yaml
	meta, err := ValidateModule(tmpDir)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// override version if passed via CLI
	if version != "" {
		meta.Version = version
	}

	// 🔹 Generate README / module.json
	if err := GenerateManifest(tmpDir, meta); err != nil {
		return fmt.Errorf("manifest failed: %w", err)
	}

	// 🔹 Init + commit
	if err := InitRepo(tmpDir, repo); err != nil {
		return fmt.Errorf("git init failed: %w", err)
	}

	// 🔹 Tag
	if err := TagVersion(tmpDir, meta.Version); err != nil {
		return fmt.Errorf("tag failed: %w", err)
	}

	// 🔹 Push
	if err := PushRepo(tmpDir); err != nil {
		return fmt.Errorf("push failed: %w", err)
	}

	// 🔹 Update local registry
	if err := UpdateRegistry(meta, repo); err != nil {
		return fmt.Errorf("registry update failed: %w", err)
	}

	return nil
}
