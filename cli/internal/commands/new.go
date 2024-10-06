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
			name          string
			addWebProject bool
		)

		if err := huh.NewInput().Title("Name").Value(&name).Run(); err != nil {
			fmt.Println(err)
			return
		}

		if err := huh.NewConfirm().Title("Add Web Project").Value(&addWebProject).Run(); err != nil {
			fmt.Println(err)
			return
		}

		if err := generate.CreateNew(name, addWebProject); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("New App %s created successfully\n", name)
	},
}
