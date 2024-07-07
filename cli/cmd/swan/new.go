package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/cli/internal"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new app",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			name string
		)

		if err := huh.NewInput().Title("Name").Value(&name).Run(); err != nil {
			panic(err)
		}

		if err := internal.CreateNew(name); err != nil {
			panic(err)
		}

		fmt.Printf("New App %s created successfully\n", name)
	},
}

func init() {
	rootCmd.AddCommand(NewCmd)
}
