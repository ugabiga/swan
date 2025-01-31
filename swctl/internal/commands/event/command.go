package event

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "make:event [path]",
	Short: "Create a new event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		fmt.Printf("Creating event at %s\n", path)

		if err := Generate(path); err != nil {
			fmt.Printf("Error while creating event: %s", err)
			return
		}

		fmt.Printf("Event %s created successfully\n", path)
	},
}
