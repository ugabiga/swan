package command

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ugabiga/swan/swctl/internal/generate"
	"github.com/ugabiga/swan/swctl/internal/utils"
)

//go:embed *.tmpl
var templateFS embed.FS

func Generate(path string) error {
	folderPath := "internal/" + path
	packageName := utils.ExtractPackageName(folderPath)

	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := generateCommandFile(folderPath, packageName); err != nil {
		return err
	}

	if err := generate.RegisterInvokeInContainer(folderPath, packageName, "SetCommand"); err != nil {
		return err
	}

	return nil
}

func generateCommandFile(folderPath, packageName string) error {
	filePath := filepath.Join(folderPath, "command.go")

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFS(templateFS, "template.tmpl")
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, map[string]string{
		"PackageName": packageName,
	}); err != nil {
		return err
	}

	return nil
}
