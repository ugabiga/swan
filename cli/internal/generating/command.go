package generating

import (
	"log"
	"os"
)

func CreateCommand(path string) {
	folderPath := "internal/" + path

	//Check if folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			panic(err)
		}
	}

	if err := createCommand(folderPath); err != nil {
		panic(err)
	}
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

	if err := registerStructToAppInvoker("./internal/config/app.go",
		fullPackageName, packageName, funcName); err != nil {
		log.Printf("Error while register struct %s", err)
		return err
	}

	return nil
}
