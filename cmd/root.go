package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sharkweb",
	Short: "Sharkweb IT Park CLI - Sharkweb Product Factory [Architect: Prathamesh Wadile]",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
