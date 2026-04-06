package generator

import (
	"path/filepath"
	"strings"
)

func GenerateModule(name string) error {

	component := strings.Title(name)

	data := TemplateData{
		Name:      name,
		Component: component,
	}

	basePath := "./" + name + "-module"

	// 🔹 Backend
	RenderTemplate(
		"internal/generator/templates/module/backend/routes.go.tmpl",
		filepath.Join(basePath, "backend/routes.go"),
		data,
	)

	// 🔹 Frontend index
	RenderTemplate(
		"internal/generator/templates/module/frontend/index.ts.tmpl",
		filepath.Join(basePath, "frontend/index.ts"),
		data,
	)

	// 🔹 Frontend page
	RenderTemplate(
		"internal/generator/templates/module/frontend/page.tsx.tmpl",
		filepath.Join(basePath, "frontend/"+component+"Page.tsx"),
		data,
	)

	// 🔹 YAML
	RenderTemplate(
		"internal/generator/templates/module/module.yaml.tmpl",
		filepath.Join(basePath, "module.yaml"),
		data,
	)

	return nil
}
