package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeDomainCmd = &cobra.Command{
	Use:   "make:domain",
	Short: "Create a new domain",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			domainName  string
			routePrefix string
		)

		if err := huh.NewInput().Title("Domain Name").Value(&domainName).Run(); err != nil {
			panic(err)
		}

		if err := huh.NewInput().Title("Route Prefix(eg: /api)").Value(&routePrefix).Run(); err != nil {
			panic(err)
		}

		generating.CreateDomain(domainName, routePrefix)

		fmt.Printf("Domain %s created successfully\n", domainName)
	},
}
