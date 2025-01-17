package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generate"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeEventCommand = &cobra.Command{
	Use:        "make:event",
	Short:      "Create a new event",
	Args:       cobra.MaximumNArgs(1),
	ArgAliases: []string{"path"},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
		)

		if len(args) > 0 {
			path = args[0]
		} else {
			if err := huh.NewInput().
				Title("Path(under internal/): ").Value(&path).Run(); err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := generate.CreateEvent(path); err != nil {
			fmt.Printf("Error while creating event: %s", err)
			return
		}

		fmt.Printf("Event %s created successfully\n", path)
	},
}
