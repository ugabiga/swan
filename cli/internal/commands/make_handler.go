package commands

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/ugabiga/swan/cli/internal/generate"

	"github.com/spf13/cobra"
)

var MakeHandlerCommand = &cobra.Command{
	Use:   "make:handler",
	Short: "Create a new domain",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlerName string
			routePrefix string
		)

		if len(args) == 2 {
			handlerName = args[0]
			routePrefix = args[1]
		} else {
			if err := huh.NewInput().Title("Handler Name").Value(&handlerName).Run(); err != nil {
				fmt.Println(err)
				return
			}

			if err := huh.NewInput().Title("Route Prefix(eg: /api)").Value(&routePrefix).Run(); err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := generate.CreateHandler(handlerName, routePrefix); err != nil {
			fmt.Printf("Error while creating domain: %s", err)
			return
		}

		fmt.Printf("Domain %s created successfully\n", handlerName)
	},
}
