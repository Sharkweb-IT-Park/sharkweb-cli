package publish

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateManifest(dir string, meta *ModuleMeta) error {

	meta.Repo = "" // will be set later

	// module.json
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dir, "module.json"), data, 0644); err != nil {
		return err
	}

	// README.md
	readme := fmt.Sprintf(`# %s

Sharkweb Module 🚀

## Version
%s

## Structure
- backend/
- frontend/
- shared/
`, meta.Name, meta.Version)

	return os.WriteFile(filepath.Join(dir, "README.md"), []byte(readme), 0644)
}
