package module

import (
	"fmt"
	"os"
	"os/exec"
)

func CloneModule(repo string, target string) error {
	fmt.Println("⬇️ Cloning:", repo)

	cmd := exec.Command("git", "clone", repo, target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
