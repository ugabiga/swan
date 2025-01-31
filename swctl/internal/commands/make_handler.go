package commands

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/ugabiga/swan/swctl/internal/generate"

	"github.com/spf13/cobra"
)

var MakeHandlerCommand = &cobra.Command{
	Use:   "make:handler",
	Short: "Create a new domain",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlerPath string
			handlerName string
			routePrefix string
		)

		if len(args) == 2 {
			handlerPath = args[0]
			handlerName = args[1]
			routePrefix = args[2]
		} else {
			if err := huh.NewInput().Title("Handler Path").Value(&handlerPath).Run(); err != nil {
				fmt.Println(err)
				return
			}

			if err := huh.NewInput().Title("Handler Name").Value(&handlerName).Run(); err != nil {
				fmt.Println(err)
				return
			}

			if err := huh.NewInput().Title("Route Prefix(eg: /api)").Value(&routePrefix).Run(); err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := generate.CreateHandler(handlerPath, handlerName, routePrefix); err != nil {
			fmt.Printf("Error while creating domain: %s", err)
			return
		}

		fmt.Printf("Domain %s created successfully\n", handlerName)
	},
}
