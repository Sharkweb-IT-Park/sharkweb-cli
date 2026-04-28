package module

import (
	"os"
	"os/exec"
)

// =========================
// 🔹 CLONE MODULE
// =========================
func CloneModule(repo string, dst string) error {

	cmd := exec.Command("git", "clone", repo, dst)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
