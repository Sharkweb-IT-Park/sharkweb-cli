package cmd

import (
	"sharkweb-cli/internal/generator"
	"sharkweb-cli/internal/utils"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate resources",
}

var generateModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generate a new module",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		utils.Step("Generating module: " + name)

		err := generator.GenerateModule(name)
		if err != nil {
			utils.Error(err.Error())
			return
		}

		utils.Success("Module generated successfully 🚀")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generateModuleCmd)
}
