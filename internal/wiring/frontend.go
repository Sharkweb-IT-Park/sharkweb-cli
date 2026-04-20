package wiring

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GenerateFrontendWiring(projectRoot string, modules []string) error {

	sort.Strings(modules)

	outputPath := filepath.Join(projectRoot, "frontend/modules/modules.gen.ts")

	var imports []string
	var registrations []string

	for _, m := range modules {

		moduleVar := toCamelCase(m) + "Module"
		moduleExport := toPascalCase(m) + "Module"

		moduleDir := filepath.Join(projectRoot, "frontend/modules", m)

		// ❌ Skip if module folder doesn't exist
		if _, err := os.Stat(moduleDir); os.IsNotExist(err) {
			fmt.Println("⚠️ Skipping missing frontend module:", m)
			continue
		}

		// ✅ NEW: Only include modules that have index.ts
		indexFile := filepath.Join(moduleDir, "index.ts")
		if _, err := os.Stat(indexFile); os.IsNotExist(err) {
			fmt.Println("⚠️ Skipping module without index.ts:", m)
			continue
		}

		importPath := fmt.Sprintf("@/modules/%s", m)

		imports = append(imports,
			fmt.Sprintf(`import { %s as %s } from "%s"`, moduleExport, moduleVar, importPath),
		)

		registrations = append(registrations, moduleVar)
	}

	code := fmt.Sprintf(`// AUTO-GENERATED FILE — DO NOT EDIT

%s

export function loadModules() {
  return [
%s
  ]
}
`,
		strings.Join(imports, "\n"),
		indent(strings.Join(registrations, ",\n"), 4),
	)

	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, []byte(code), 0644); err != nil {
		return err
	}

	fmt.Println("✅ Frontend wiring generated:", outputPath)

	return nil
}
func GenerateNextRoutes(projectRoot string, modules []string) error {

	appDir := filepath.Join(projectRoot, "frontend/app")

	for _, m := range modules {

		srcPage := filepath.Join(
			projectRoot,
			"frontend/modules",
			m,
			"pages/page.tsx",
		)

		destDir := filepath.Join(appDir, m)
		destPage := filepath.Join(destDir, "page.tsx")

		// skip if module has no page
		if _, err := os.Stat(srcPage); os.IsNotExist(err) {
			continue
		}

		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return err
		}

		content := fmt.Sprintf(`export { default } from "@/modules/%s/pages/page"`, m)

		if err := os.WriteFile(destPage, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}
