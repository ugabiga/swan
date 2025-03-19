package entdb

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "ent:new [name]",
	Short: "Create a new ent schema",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Println("Creating new ent schema")
		err := runEntCommand("new", name)
		if err != nil {
			fmt.Println("failed to create new ent schema", "error", err)
		}

		fmt.Printf("Command %s created successfully\n", name)
	},
}

var GenerateCmd = &cobra.Command{
	Use:   "ent:generate",
	Short: "Generate ent code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating ent code")
		err := runEntCommand("generate")
		if err != nil {
			fmt.Println("failed to generate ent code", "error", err)
		}
	},
}

func runEntCommand(cmd string, args ...string) error {
	path := "./internal/ent/schema"
	entCmd := exec.Command("go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", cmd)

	switch cmd {
	case "new":
		entCmd.Args = append(entCmd.Args, "--target", path, args[0])
	case "generate":
		entCmd.Args = append(entCmd.Args, path)
	}

	output, err := entCmd.CombinedOutput()
	if err != nil {
		fmt.Println("failed to create migration:", err)
		fmt.Println("command output:", string(output))
		return err
	}

	return nil
}
