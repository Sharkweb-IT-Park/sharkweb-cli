package project

import (
	"fmt"
	"os/exec"

	"sharkweb-cli/internal/module"
)

func SetupFullbase(projectRoot string) error {

	fmt.Println("🎨 Setting up frontend...")

	return module.CloneModule(
		"https://github.com/Sharkweb-IT-Park/sharkweb-mvp-base.git",
		projectRoot,
	)
}

func InstallFrontendDeps(projectRoot string) error {
	fmt.Println("📦 Installing frontend dependencies...")

	cmd := exec.Command("npm", "install")
	cmd.Dir = projectRoot + "/frontend"

	return cmd.Run()
}

func InstallBackendDeps(projectRoot string) error {

	fmt.Println("📦 Installing backend dependencies...")

	cmd := exec.Command("go", "get", "github.com/gin-gonic/gin")
	cmd.Dir = projectRoot + "/backend"

	return cmd.Run()
}
