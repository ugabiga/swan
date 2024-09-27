package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeCommandCommand = &cobra.Command{
	Use:   "make:command",
	Short: "Create a new command",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			name string
		)

		if err := huh.NewInput().
			Title("Path").
			Value(&path).
			Run(); err != nil {
			panic(err)
		}

		generating.CreateCommand(path, name)

		fmt.Printf("Struct %s created successfully\n", name)
	},
}
