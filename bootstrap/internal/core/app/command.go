package app

import (
	"log"

	"github.com/spf13/cobra"
)

func RunCommand(rootCmd *cobra.Command) {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Application CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return rootCmd
}
