package cmd

import (
	"fmt"
	"os"

	"sharkweb-cli/internal/project"
	"sharkweb-cli/internal/utils"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Sharkweb app",
}

var createAppCmd = &cobra.Command{
	Use:   "app [name]",
	Short: "Create a new Sharkweb app",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		appName := args[0]

		// =========================
		// 🛑 Prevent overwrite
		// =========================
		if _, err := os.Stat(appName); err == nil {
			utils.Error("Folder already exists: " + appName)
			return
		}

		utils.Step("Creating app: " + appName)

		// =========================
		// 🔥 FULL PROJECT SETUP
		// =========================
		err := project.Setup(appName)
		if err != nil {
			utils.Error("Failed to create app: " + err.Error())
			return
		}

		// =========================
		// 🎉 SUCCESS OUTPUT
		// =========================
		fmt.Println()
		utils.Success("Sharkweb app created successfully 🚀")
		fmt.Println()
		fmt.Println("👉 Next steps:")
		fmt.Println("cd", appName)
		fmt.Println("sharkweb dev")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createAppCmd)
}
