package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "snacks",
	Short:   "util collection",
	Long:    "util collection",
	Example: "snacks SUBCOMMAND",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
