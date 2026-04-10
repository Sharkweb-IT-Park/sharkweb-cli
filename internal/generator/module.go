package generator

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"sharkweb-cli/internal/utils"
	"sharkweb-cli/internal/wiring"
)

type TemplateData struct {
	Name        string
	Component   string
	HasBackend  bool
	HasFrontend bool
	ModulePath  string
}

func GenerateModule(name string) error {

	component := ToPascalCase(name)

	// ✅ Get project root
	projectRoot, err := utils.ValidateProjectRoot()
	if err != nil {
		return err
	}

	goMod, err := GetGoModuleName(projectRoot + "/backend")
	if err != nil {
		return err
	}

	data := TemplateData{
		Name:        name,
		Component:   component,
		HasBackend:  true,
		HasFrontend: false,
		ModulePath:  goMod, // ✅ REQUIRED
	}

	// ✅ Paths
	backendBasePath := filepath.Join(projectRoot, "backend", "modules", name)
	frontendBasePath := filepath.Join(projectRoot, "frontend", "modules", name)

	// 🔥 isolate template root
	subFS, err := fs.Sub(templateFS, "templates/module/full")
	if err != nil {
		return err
	}

	// ✅ Generate module files
	if err := GenerateFromEmbedFS(subFS, backendBasePath, frontendBasePath, data); err != nil {
		return err
	}

	// 🔥 NEW: Auto Wiring System

	// 1️⃣ Detect all modules
	modules, err := GetModules(projectRoot)
	if err != nil {
		return err
	}

	// 2️⃣ Backend wiring
	if err := wiring.GenerateBackendWiring(projectRoot, modules); err != nil {
		return err
	}

	// 3️⃣ Frontend wiring
	if err := wiring.GenerateFrontendWiring(projectRoot, modules); err != nil {
		return err
	}

	return nil
}
func GenerateFromEmbedFS(fsys fs.FS, backendDir string, frontendDir string, data TemplateData) error {

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		// 🔥 Replace dynamic names
		targetPath := path
		targetPath = strings.ReplaceAll(targetPath, "{{.Name}}", data.Name)
		targetPath = strings.ReplaceAll(targetPath, "{{.Component}}", data.Component)

		// 🔥 Decide where to generate
		var finalPath string

		if strings.HasPrefix(path, "backend") {
			finalPath = filepath.Join(backendDir, strings.TrimPrefix(targetPath, "backend"))
		} else if strings.HasPrefix(path, "frontend") {
			finalPath = filepath.Join(frontendDir, strings.TrimPrefix(targetPath, "frontend"))
		} else {
			// shared or root files → put inside backend module (or customize later)
			finalPath = filepath.Join(backendDir, targetPath)
		}

		if d.IsDir() {
			return os.MkdirAll(finalPath, os.ModePerm)
		}

		// Remove .tmpl extension
		finalPath = strings.TrimSuffix(finalPath, ".tmpl")

		return RenderTemplateFromFS(path, finalPath, data, fsys)
	})
}
