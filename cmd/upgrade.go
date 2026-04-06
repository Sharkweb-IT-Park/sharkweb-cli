package cmd

import (
	"fmt"

	"sharkweb-cli/internal/config"
	"sharkweb-cli/internal/module"
	"sharkweb-cli/internal/utils"

	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade resources of project",
}

var upgradeModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Upgrade module to latest version",
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
		// 🔹 Check installed
		// =========================
		if !config.IsModuleInstalled(cfg, moduleName) {
			utils.Error(fmt.Sprintf("%s is not installed", moduleName))
			return
		}

		utils.Step("Upgrading module: " + moduleName)

		// =========================
		// 🔥 UPGRADE
		// =========================
		err = module.UpgradeModule(moduleName, projectRoot)
		if err != nil {
			utils.Error("Upgrade failed: " + err.Error())
			return
		}

		utils.Success("Module upgraded successfully 🚀")
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
	upgradeCmd.AddCommand(upgradeModuleCmd)
}
