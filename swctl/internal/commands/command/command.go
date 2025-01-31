package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "make:command [path]",
	Short: "Create a new command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		fmt.Printf("Creating command at %s\n", path)

		if err := Generate(path); err != nil {
			fmt.Printf("Error while creating command: %s", err)
			return
		}

		fmt.Printf("Command %s created successfully\n", path)
	},
}
