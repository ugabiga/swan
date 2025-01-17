package commands

import (
	"fmt"

	"github.com/ugabiga/swan/cli/internal/generate"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new app",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			name string
		)

		if err := huh.NewInput().Title("Name v2").Value(&name).Run(); err != nil {
			fmt.Println(err)
			return
		}

		if err := generate.CreateNew(name); err != nil {
			fmt.Printf("Error while creating app: %s", err)
			return
		}

		fmt.Printf("New App %s created successfully\n", name)
	},
}
