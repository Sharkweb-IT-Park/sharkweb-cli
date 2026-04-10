package wiring

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GenerateFrontendWiring(projectRoot string, modules []string) error {

	// 🔥 Ensure stable output
	sort.Strings(modules)

	outputPath := filepath.Join(projectRoot, "frontend/modules/modules.gen.ts")

	var imports []string
	var registrations []string

	for _, m := range modules {

		moduleVar := toCamelCase(m) + "Module"
		moduleExport := toPascalCase(m) + "Module"

		importPath := fmt.Sprintf("@/modules/%s/frontend", m)

		// 🔥 Check if module exists
		moduleDir := filepath.Join(projectRoot, "frontend/modules", m)
		if _, err := os.Stat(moduleDir); os.IsNotExist(err) {
			fmt.Println("⚠️ Skipping missing frontend module:", m)
			continue
		}

		imports = append(imports,
			fmt.Sprintf(`import { %s as %s } from "%s"`, moduleExport, moduleVar, importPath),
		)

		registrations = append(registrations, moduleVar)
	}

	code := fmt.Sprintf(`// AUTO-GENERATED FILE — DO NOT EDIT

import { AppModule } from "@/core/module"

%s

export function loadModules(): AppModule[] {
  return [
%s
  ]
}
`,
		strings.Join(imports, "\n"),
		indent(strings.Join(registrations, ",\n"), 4),
	)

	// 🔥 Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return err
	}

	// 🔥 Write file
	if err := os.WriteFile(outputPath, []byte(code), 0644); err != nil {
		return err
	}

	fmt.Println("✅ Frontend wiring generated:", outputPath)

	return nil
}
