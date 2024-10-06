package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/cli/internal/generate"
)

var MakeDBClient = &cobra.Command{
	Use:   "init:ent",
	Short: "Initialize Ent ORM",
	Run: func(cmd *cobra.Command, args []string) {
		err := generate.InitializeEntORM()
		if err != nil {
			switch {
			case errors.Is(err, generate.ErrEntNotInitialized):
				fmt.Printf("Ent not initialized: Please run 'make ent-new and make ent-gen' first\n")
				return
			default:
				fmt.Println(err)
				return
			}

		}

		fmt.Printf("Successfully initialized ORM\n")
	},
}
