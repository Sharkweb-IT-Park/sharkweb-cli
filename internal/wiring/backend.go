package wiring

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// ✅ Embed templates
//
//go:embed templates/**
var templateFS embed.FS

type Import struct {
	Alias string
	Path  string
}

type WiringData struct {
	ModulePath string // ✅ REQUIRED (this fixes your error)
	Imports    []Import
	Modules    []string
}

func GenerateBackendWiring(projectRoot string, modules []string) error {

	backendRoot := filepath.Join(projectRoot, "backend")
	outputPath := filepath.Join(backendRoot, "modules", "modules.gen.go")

	modulePath, err := GetBackendModulePath(projectRoot)
	if err != nil {
		return err
	}

	var imports []Import
	var registrations []string

	for _, m := range modules {

		imports = append(imports, Import{
			Alias: m,
			Path:  fmt.Sprintf("%s/modules/%s", modulePath, m),
		})

		registrations = append(registrations,
			fmt.Sprintf("%s.NewModule()", m),
		)
	}

	data := WiringData{
		ModulePath: modulePath,
		Imports:    imports,
		Modules:    registrations,
	}

	// 🔥 LOAD FROM EMBED (THIS FIXES EVERYTHING)
	tpl, err := template.ParseFS(
		WiringTemplates,
		"templates/wiring/backend/wiring.go.tpl",
	)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tpl.Execute(&buf, data); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}
