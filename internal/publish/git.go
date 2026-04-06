package publish

import (
	"os/exec"
)

func TagVersion(modulePath string, version string) error {

	cmd := exec.Command("git", "tag", "v"+version)
	cmd.Dir = modulePath

	return cmd.Run()
}

func PushRepo(modulePath string) error {

	cmd := exec.Command("git", "push", "--tags")
	cmd.Dir = modulePath

	return cmd.Run()
}
