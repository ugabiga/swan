package commands

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/ugabiga/swan/cli/internal/generating"

	"github.com/spf13/cobra"
)

var MakeCommandCommand = &cobra.Command{
	Use:   "make:command",
	Short: "Create a new command",
	Args:  cobra.MaximumNArgs(1),
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

		generating.CreateCommand(path)

		fmt.Printf("Command created successfully at %s\n", path)
	},
}
