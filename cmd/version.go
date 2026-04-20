package cmd

import (
	"fmt"

	"sharkweb-cli/internal/version"

	"github.com/spf13/cobra"
)

var short bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			fmt.Println(version.Short())
			return
		}
		fmt.Println(version.Info())
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&short, "short", "s", false, "Short version")
	rootCmd.AddCommand(versionCmd)
}
