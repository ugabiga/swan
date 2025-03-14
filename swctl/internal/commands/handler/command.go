package handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "make:handler [file_path_under_internal] [route_prefix] [handler_name]",
	Short:   "Create a new handler, file_path starts from internal",
	Example: "swctl make:handler app/handlers /api/v1 users",
	Args:    cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlerFilePath  string
			routePrefix      string
			handlerRouteName string
		)

		if len(args) == 0 {
			cmd.Help()
			return
		}

		handlerFilePath = args[0]
		routePrefix = args[1]
		handlerRouteName = args[2]

		if err := Generate(handlerFilePath, routePrefix, handlerRouteName); err != nil {
			fmt.Printf("Error while creating handler: %s", err)
			return
		}

		fmt.Printf("Handler %s created successfully\n", handlerRouteName)
	},
}
