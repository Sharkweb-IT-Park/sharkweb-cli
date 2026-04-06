package module

import (
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies a file from src → dst
func CopyFile(src, dst string) error {

	// 1. Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 2. Ensure destination directory exists
	err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}

	// 3. Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 4. Copy contents
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 5. Preserve file permissions (optional but recommended)
	info, err := os.Stat(src)
	if err == nil {
		_ = os.Chmod(dst, info.Mode())
	}

	return nil
}
