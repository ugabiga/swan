package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeEventCommand = &cobra.Command{
	Use:   "make:event",
	Short: "Create a new event",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			name string
		)

		if err := huh.NewInput().
			Title("Path(under internal/): ").
			Value(&path).
			Run(); err != nil {
			panic(err)
		}

		generating.CreateEvent(path)

		fmt.Printf("Struct %s created successfully\n", name)
	},
}
