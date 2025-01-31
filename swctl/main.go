package main

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/command"
	"github.com/ugabiga/swan/swctl/internal/commands"
)

var rootCmd = &cobra.Command{
	Use: "swan",
}

func main() {
	for _, cmd := range commands.Commands() {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.AddCommand(command.CreateCmds())

	_ = rootCmd.Execute()
}
