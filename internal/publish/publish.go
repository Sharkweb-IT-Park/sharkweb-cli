package publish

import (
	"fmt"
)

func PublishModule(modulePath string, repo string) error {

	fmt.Println("🚀 Publishing module...")

	// 1. Validate module
	meta, err := ValidateModule(modulePath)
	if err != nil {
		return err
	}

	// 2. Git tag
	err = TagVersion(modulePath, meta.Version)
	if err != nil {
		return err
	}

	// 3. Push repo
	err = PushRepo(modulePath)
	if err != nil {
		return err
	}

	// 4. Update registry
	err = UpdateRegistry(meta, repo)
	if err != nil {
		return err
	}

	fmt.Println("✅ Module published:", meta.Name)

	return nil
}
