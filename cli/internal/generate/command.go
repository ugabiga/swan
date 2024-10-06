package generate

import (
	"github.com/ugabiga/swan/cli/internal/utils"
	"log"
	"os"
)

func CreateCommand(path string) error {
	folderPath := "internal/" + path

	//Check if folder exists
	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := createCommand(folderPath); err != nil {
		return err
	}

	return nil
}

func createCommand(folderPath string) error {
	fileName := "command"
	filePath := folderPath + "/" + fileName + ".go"
	fullPackageName := folderPath
	packageName := extractPackageName(folderPath)
	funcName := "InvokeSetCommands"

	template := `package ` + packageName + `

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/core"
)

func InvokeSetCommands(
	command *core.Command,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "cmd",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
			},
		},
	)
}

`
	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		log.Printf("Error while creating struct: %s", err)
		return err
	}

	if err := registerToInvoker(
		"./internal/config/commands.go",
		fullPackageName,
		packageName,
		funcName,
	); err != nil {
		log.Printf("Error while register struct %s", err)
		return err
	}

	return nil
}
