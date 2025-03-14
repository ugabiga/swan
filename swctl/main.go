package main

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/commands"
)

var rootCmd = &cobra.Command{
	Use: "swctl",
}

func main() {
	for _, cmd := range commands.Commands() {
		rootCmd.AddCommand(cmd)
	}

	_ = rootCmd.Execute()
}
