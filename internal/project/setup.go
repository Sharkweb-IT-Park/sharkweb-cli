package project

import "fmt"

func Setup(name string) error {

	fmt.Println("🚀 Creating Sharkweb app:", name)

	root, err := InitProject(name)
	if err != nil {
		return err
	}

	err = SetupBackend(root)
	if err != nil {
		return err
	}

	err = SetupFrontend(root)
	if err != nil {
		return err
	}

	err = InstallFrontendDeps(root)
	if err != nil {
		return err
	}

	err = CreateConfig(root)
	if err != nil {
		return err
	}

	fmt.Println("✅ Project created successfully")

	return nil
}
