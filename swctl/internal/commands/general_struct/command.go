package general_struct

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "make:struct [path] [name]",
	Short: "Create a new struct",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		name := args[1]

		fmt.Printf("Creating struct at %s with name %s\n", path, name)

		if err := Generate(path, name); err != nil {
			fmt.Printf("Error while creating struct: %s", err)
			return
		}

		fmt.Printf("Struct %s created successfully\n", path)
	},
}
