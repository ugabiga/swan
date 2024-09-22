package main

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/cli/internal/commands"
)

var rootCmd = &cobra.Command{
	Use: "swan",
}

func main() {
	for _, cmd := range commands.Commands() {
		rootCmd.AddCommand(cmd)
	}

	_ = rootCmd.Execute()
}
