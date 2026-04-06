package cmd

import (
	"sharkweb-cli/internal/dev"
	"sharkweb-cli/internal/utils"

	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run full development environment",

	Run: func(cmd *cobra.Command, args []string) {

		projectRoot, err := utils.ValidateProjectRoot()

		if err != nil {
			utils.Error(err.Error())
			return
		}

		utils.Step("Starting development environment...")

		err = dev.RunDev(projectRoot)

		if err != nil {
			utils.Error("Dev failed: " + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
