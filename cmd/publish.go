package cmd

import (
	"sharkweb-cli/internal/publish"
	"sharkweb-cli/internal/utils"

	"github.com/spf13/cobra"
)

var repo string
var versionmodule string

// =========================
// 🔹 PARENT COMMAND
// =========================
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish resources",
}

// =========================
// 🔹 CHILD COMMAND
// =========================
var publishModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Publish module",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		moduleName := args[0]

		if repo == "" {
			utils.Error("repo is required (--repo)")
			return
		}

		utils.Step("Publishing module...")

		err := publish.PublishModule(moduleName, repo, versionmodule)
		if err != nil {
			utils.Error("Publish failed: " + err.Error())
			return
		}

		utils.Success("Module published 🚀")
	},
}

// =========================
// 🔹 INIT
// =========================
func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.AddCommand(publishModuleCmd)

	// flags
	publishModuleCmd.Flags().StringVar(&repo, "repo", "", "GitHub repo URL")
	publishModuleCmd.Flags().StringVar(&versionmodule, "version", "", "Module version (optional, overrides module.yaml)")
}
