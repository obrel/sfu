package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "obrel-sfu",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		SilenceUsage: true,
	}
)

func Execute() {
	rootCmd.AddCommand(serverCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
