package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	Name      string
	Component string
}

func RenderTemplate(templatePath, outputPath string, data TemplateData) error {

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
