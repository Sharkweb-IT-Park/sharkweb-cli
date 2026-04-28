package cmd

import (
	"fmt"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/module"
	"sharkweb-cli/internal/utils"

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

		projectRoot, err := utils.ValidateProjectRoot()
		if err != nil {
			utils.Error(err.Error())
			return
		}

		cfg, err := config.Load(projectRoot)
		if err != nil {
			utils.Error("Failed to load config")
			return
		}

		if config.IsModuleInstalled(cfg, moduleName) {
			utils.Info(fmt.Sprintf("%s already installed", moduleName))
			return
		}

		utils.Step("Installing module: " + moduleName)

		// 🔥 USE CORRECT FLOW
		err = module.AddModule(moduleName, projectRoot)
		if err != nil {
			utils.Error("Installation failed: " + err.Error())
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
