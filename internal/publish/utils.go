package publish

import (
	"os"
	"path/filepath"
)

// =========================
// 🔹 CHECK PATH EXISTS
// =========================
func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// =========================
// 🔹 COPY FILE
// =========================
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// ensure parent dir exists
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(dst, data, 0644)
}

// =========================
// 🔹 COPY DIRECTORY
// =========================
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, os.ModePerm)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(target, data, info.Mode())
	})
}

func copyIfExists(src, dst string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return nil
	}
	return copyDir(src, dst)
}
