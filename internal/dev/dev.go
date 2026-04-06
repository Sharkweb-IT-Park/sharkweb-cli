package dev

func RunDev(projectRoot string) error {

	supervisor := &Supervisor{}

	// 🔹 Frontend
	supervisor.Add(&Process{
		Name:    "frontend",
		Command: "npm",
		Args:    []string{"run", "dev"},
		Dir:     projectRoot + "/frontend",
	})

	// 🔹 Backend (HOT RELOAD via air)
	supervisor.Add(&Process{
		Name:    "backend",
		Command: "air",
		Args:    []string{},
		Dir:     projectRoot + "/backend",
	})

	// 🔹 Start everything
	supervisor.StartAll()

	return nil
}
