package cmd

import (
	"fmt"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/module"
	"sharkweb-cli/internal/utils"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove resources from project",
}

var removeModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Remove module from project",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		moduleName := args[0]

		// =========================
		// 🔹 Validate project
		// =========================
		projectRoot, err := utils.ValidateProjectRoot()
		if err != nil {
			utils.Error(err.Error())
			return
		}

		// =========================
		// 🔹 Load config
		// =========================
		cfg, err := config.Load(projectRoot)
		if err != nil {
			utils.Error("Failed to load config")
			return
		}

		// =========================
		// 🔹 Check exists
		// =========================
		if !config.IsModuleInstalled(cfg, moduleName) {
			utils.Error(fmt.Sprintf("%s is not installed", moduleName))
			return
		}

		utils.Step("Removing module: " + moduleName)

		// =========================
		// 🔥 REMOVE MODULE
		// =========================
		err = module.RemoveModule(moduleName, projectRoot)
		if err != nil {
			utils.Error("Remove failed: " + err.Error())
			return
		}

		utils.Success("Module removed successfully 🚀")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.AddCommand(removeModuleCmd)
}
