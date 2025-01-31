package generate

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ugabiga/swan/swctl/internal/utils"
)

func CreateCommand(path string) error {
	folderPath := "internal/" + path
	fileName := "command"
	filePath := folderPath + "/" + fileName + ".go"
	invokerFilePath := CommandPath
	packageName := extractPackageName(folderPath)
	funcName := "SetCommands"

	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := commandTemplate(
		filePath,
		packageName,
		funcName,
	); err != nil {
		return err
	}

	if err := registerToInvoker(
		invokerFilePath,
		folderPath,
		packageName,
		funcName,
	); err != nil {
		fmt.Printf("Error while registering command to invoker: %s", err)
		return err
	}

	return nil
}

func commandTemplate(filePath, packageName, funcName string) error {
	type CommandTemplateData struct {
		PackageName string
		FuncName    string
	}

	tmplData := CommandTemplateData{
		PackageName: packageName,
		FuncName:    funcName,
	}

	tmpl, err := template.New("command").Parse(`package {{.PackageName}}

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/core"
)

func {{.FuncName}}(
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
`)
	if err != nil {
		fmt.Printf("Error while parsing template: %s", err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error while creating file: %s", err)
		return err
	}
	defer file.Close()

	if err = tmpl.Execute(file, tmplData); err != nil {
		fmt.Printf("Error while executing template: %s", err)
		return err
	}

	return nil
}
