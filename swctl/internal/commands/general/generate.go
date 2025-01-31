package general

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

func Generate(path, name string) error {
	folderPath := "internal/" + path
	packageName := utils.ExtractPackageName(folderPath)
	funcName := "New" + name

	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := generateStructFile(folderPath, packageName, name); err != nil {
		return err
	}

	if err := generate.RegisterStructToContainer(folderPath, packageName, funcName); err != nil {
		return err
	}

	return nil
}

func generateStructFile(folderPath, packageName, structName string) error {
	snakeName := utils.ConvertToSnakeCase(structName)
	filePath := filepath.Join(folderPath, snakeName+".go")

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
		"StructName":  structName,
	}); err != nil {
		return err
	}

	return nil
}
