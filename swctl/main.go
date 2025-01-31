package main

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/commands"
	"github.com/ugabiga/swan/swctl/internal/commands_v2"
)

var rootCmd = &cobra.Command{
	Use: "swan",
}

func main() {
	for _, cmd := range commands.Commands() {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.AddCommand(commands_v2.Commands()...)

	_ = rootCmd.Execute()
}
