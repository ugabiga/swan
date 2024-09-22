package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new app",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			name          string
			addWebProject bool
		)

		if err := huh.NewInput().Title("Name").Value(&name).Run(); err != nil {
			panic(err)
		}

		if err := huh.NewConfirm().Title("Add Web Project").Value(&addWebProject).Run(); err != nil {
			panic(err)
		}

		if err := generating.CreateNew(name, addWebProject); err != nil {
			panic(err)
		}

		fmt.Printf("New App %s created successfully\n", name)
	},
}
