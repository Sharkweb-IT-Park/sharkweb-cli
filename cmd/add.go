package cmd

import (
	"fmt"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/module"
	"sharkweb-cli/internal/utils"
	"sharkweb-cli/internal/wiring"

	"github.com/spf13/cobra"
)

// =========================
// 🔹 PARENT COMMAND
// =========================
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add resources to project",
}

// =========================
// 🔹 CHILD COMMAND
// =========================
var addModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Add module to project",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		moduleName := args[0]

		// =========================
		// 🔹 1. Validate project
		// =========================
		projectRoot, err := utils.ValidateProjectRoot()
		if err != nil {
			utils.Error(err.Error())
			return
		}

		// =========================
		// 🔹 2. Load config
		// =========================
		cfg, err := config.Load(projectRoot)
		if err != nil {
			utils.Error("Failed to load config")
			return
		}

		// =========================
		// 🔹 3. Prevent duplicate
		// =========================
		if config.IsModuleInstalled(cfg, moduleName) {
			utils.Info(fmt.Sprintf("%s already installed", moduleName))
			return
		}

		utils.Step("Installing module: " + moduleName)

		// =========================
		// 🔹 4. Install module
		// =========================
		installed := make(map[string]bool)

		err = module.InstallModule(moduleName, projectRoot, installed)
		if err != nil {
			utils.Error("Installation failed: " + err.Error())
			return
		}

		// =========================
		// 🔥 Reload config (source of truth)
		// =========================
		cfg, err = config.Load(projectRoot)
		if err != nil {
			utils.Error("Failed to reload config")
			return
		}

		utils.Step("Generating wiring...")

		// =========================
		// 🔥 AUTO-WIRING
		// =========================
		err = wiring.GenerateWiring(projectRoot, cfg.Modules)
		if err != nil {
			utils.Error("Wiring failed: " + err.Error())
			return
		}

		utils.Success("Module installed & wired successfully 🚀")
	},
}

// =========================
// 🔹 INIT
// =========================
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addModuleCmd)
}
