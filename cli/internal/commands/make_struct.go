package commands

import (
	"fmt"
	"github.com/ugabiga/swan/cli/internal/generate"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var MakeStruct = &cobra.Command{
	Use:   "make:struct",
	Short: "Create a new struct",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			name string
		)
		if len(args) == 2 {
			path = args[0]
			name = args[1]
		} else {
			if err := huh.NewInput().Title("Path(under internal/): ").Value(&path).Run(); err != nil {
				fmt.Println(err)
				return
			}
			if err := huh.NewInput().Title("Name").Value(&name).Run(); err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := generate.CreateStruct(path, name); err != nil {
			fmt.Printf("Error while creating struct: %s", err)
			return
		}

		fmt.Printf("Struct %s created successfully\n", name)
	},
}
