package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/cli/internal/generating"
)

var MakeDBClient = &cobra.Command{
	Use:   "make:init-orm",
	Short: "Create DB client, add migration setup to .env file and add Makefile commands",
	Run: func(cmd *cobra.Command, args []string) {
		err := generating.InitializeORM()
		if err != nil {
			switch {
			case errors.Is(err, generating.ErrEntNotInitialized):
				fmt.Printf("Ent not initialized: Please run 'make ent-new and make ent-gen' first\n")
				return
			default:
				panic(err)
			}

		}

		fmt.Printf("Successfully initialized ORM\n")
	},
}
