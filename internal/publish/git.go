package publish

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git %v failed: %w", args, err)
	}
	return nil
}

func InitRepo(dir string, repo string) error {

	os.RemoveAll(filepath.Join(dir, ".git"))

	if err := runGit(dir, "init"); err != nil {
		return err
	}

	if err := runGit(dir, "add", "."); err != nil {
		return err
	}

	if err := runGit(dir, "commit", "-m", "Initial commit"); err != nil {
		return err
	}

	if err := runGit(dir, "branch", "-M", "main"); err != nil {
		return err
	}

	return runGit(dir, "remote", "add", "origin", repo)
}

func TagVersion(dir string, version string) error {
	tag := "v" + version
	return runGit(dir, "tag", "-a", tag, "-m", "Release "+tag)
}

func PushRepo(dir string) error {

	if err := runGit(dir, "push", "-u", "origin", "main"); err != nil {
		return err
	}

	return runGit(dir, "push", "--tags")
}
