package wiring

func GenerateWiring(projectRoot string, modules []string) error {

	// generate backend wiring
	err := GenerateBackendWiring(projectRoot, modules)
	if err != nil {
		return err
	}

	// generate frontend wiring
	err = GenerateFrontendWiring(projectRoot, modules)
	if err != nil {
		return err
	}

	return nil
}
