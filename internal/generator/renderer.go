package generator

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

func RenderTemplateFromFS(
	fsPath string,
	outputPath string,
	data TemplateData,
	fsys fs.FS,
) error {

	// 🔥 Read template file
	content, err := fs.ReadFile(fsys, fsPath)
	if err != nil {
		return fmt.Errorf("read template error (%s): %w", fsPath, err)
	}

	// 🔥 Strict template parsing (fails if missing fields)
	tmpl, err := template.
		New(filepath.Base(fsPath)).
		Option("missingkey=error").
		Parse(string(content))
	if err != nil {
		return fmt.Errorf("template parse error (%s): %w", fsPath, err)
	}

	var buf bytes.Buffer

	// 🔥 Execute template
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("template execute error (%s): %w", fsPath, err)
	}

	// 🔥 Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return fmt.Errorf("mkdir error (%s): %w", outputPath, err)
	}

	// 🔥 Prevent overwrite (VERY IMPORTANT)
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Println("⚠️ Skipping existing file:", outputPath)
		return nil
	}

	// 🔥 Write file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("write file error (%s): %w", outputPath, err)
	}

	// ✅ Debug log (optional but useful)
	fmt.Println("✅ Generated:", outputPath)

	return nil
}
