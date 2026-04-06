package project

import (
	"fmt"
	"os/exec"

	"sharkweb-cli/internal/module"
)

func SetupFrontend(projectRoot string) error {

	fmt.Println("🎨 Setting up frontend...")

	return module.CloneModule(
		"https://github.com/sharkweb/nextjs-base",
		projectRoot+"/frontend",
	)
}

func InstallFrontendDeps(projectRoot string) error {
	fmt.Println("📦 Installing frontend dependencies...")

	cmd := exec.Command("npm", "install")
	cmd.Dir = projectRoot + "/frontend"

	return cmd.Run()
}
