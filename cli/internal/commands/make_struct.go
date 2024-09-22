package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeStruct = &cobra.Command{
	Use:   "make:struct",
	Short: "Create a new struct",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			name string
		)

		if err := huh.NewInput().Title("Path").Value(&path).Run(); err != nil {
			panic(err)
		}

		if err := huh.NewInput().Title("Name").Value(&name).Run(); err != nil {
			panic(err)
		}

		generating.CreateStruct(path, name)

		fmt.Printf("Struct %s created successfully\n", name)
	},
}
