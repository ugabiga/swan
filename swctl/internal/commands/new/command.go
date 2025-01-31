package new

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if err := Generate(name); err != nil {
			fmt.Printf("Error while creating app: %s", err)
			return
		}

		fmt.Printf("New App %s created successfully\n", name)
	},
}
