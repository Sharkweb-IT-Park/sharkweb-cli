package wiring

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/**
var templateFS embed.FS

type Import struct {
	Alias string
	Path  string
}

type WiringData struct {
	ModulePath string
	Imports    []Import
	Modules    []string
}

// =========================
// 🔹 GENERATE BACKEND WIRING
// =========================
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

		// 🔥 STANDARD CONTRACT
		registrations = append(registrations,
			fmt.Sprintf("%s.NewModule()", m),
		)
	}

	data := WiringData{
		ModulePath: modulePath,
		Imports:    imports,
		Modules:    registrations,
	}

	tpl, err := template.ParseFS(
		templateFS,
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

	fmt.Println("✅ Backend wiring generated:", outputPath)

	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}
