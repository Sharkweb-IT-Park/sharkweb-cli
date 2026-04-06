package project

import (
	"fmt"
	"os/exec"
)

func SetupBackend(projectRoot string) error {

	backendPath := projectRoot + "/backend"

	fmt.Println("⚙️ Initializing Go module...")

	cmd := exec.Command("go", "mod", "init", "backend")
	cmd.Dir = backendPath

	err := cmd.Run()
	if err != nil {
		return err
	}

	// install gin
	fmt.Println("📦 Installing backend dependencies...")

	cmd = exec.Command("go", "get", "github.com/gin-gonic/gin")
	cmd.Dir = backendPath

	return cmd.Run()
}
