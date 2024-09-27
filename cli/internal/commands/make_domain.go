package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeHandlerCommand = &cobra.Command{
	Use:   "make:handler",
	Short: "Create a new domain",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlerName string
			routePrefix string
		)

		if err := huh.NewInput().
			Title("Handler Name").
			Value(&handlerName).
			Run(); err != nil {
			panic(err)
		}

		if err := huh.NewInput().
			Title("Route Prefix(eg: /api)").
			Value(&routePrefix).
			Run(); err != nil {
			panic(err)
		}

		generating.CreateHandler(handlerName, routePrefix)

		fmt.Printf("Domain %s created successfully\n", handlerName)
	},
}
