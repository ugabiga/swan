package handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "make:handler [route-prefix] [path] [name]",
	Short: "Create a new handler",
	Args:  cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlerPath string
			handlerName string
			routePrefix string
		)

		routePrefix = args[0]
		handlerPath = args[1]
		handlerName = args[2]

		if err := Generate(routePrefix, handlerPath, handlerName); err != nil {
			fmt.Printf("Error while creating handler: %s", err)
			return
		}

		fmt.Printf("Handler %s created successfully\n", handlerName)
	},
}
