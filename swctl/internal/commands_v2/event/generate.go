package event

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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

	if err := generateEventFile(folderPath, packageName); err != nil {
		return err
	}

	if err := registerInvokeInContainer(folderPath, packageName, "SetEvent"); err != nil {
		return err
	}

	return nil
}

func generateEventFile(folderPath, packageName string) error {
	filePath := filepath.Join(folderPath, "event.go")

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

const (
	AppRootPath   = "./internal/app"
	ContainerPath = AppRootPath + "/container.go"
)

func registerInvokeInContainer(folderPath, structName, funcName string) error {
	containerPath := ContainerPath
	invokeFunc := "fx.Invoke"

	bytes, err := os.ReadFile(containerPath)
	if err != nil {
		return err
	}
	containerFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	packagePath := moduleName + "/" + folderPath
	log.Printf("packagePath: %s", packagePath)
	containerFileContents = strings.ReplaceAll(containerFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(containerFileContents, invokeFunc+"(") {
		containerFileContents = strings.Replace(containerFileContents, invokeFunc+
			"(", invokeFunc+
			"(\n\t\t\t"+structName+"."+funcName+",", 1)

		if err = os.WriteFile(containerPath, []byte(containerFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}
