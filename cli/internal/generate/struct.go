package generate

import (
	"github.com/ugabiga/swan/cli/internal/utils"
	"log"
	"os"
	"strings"
)

func CreateStruct(path, name string) error {
	folderPath := "internal/" + path

	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := createStruct(folderPath, name); err != nil {
		return err
	}

	return nil
}

func createStruct(folderPath, structName string) error {
	fileName := strings.ToLower(structName)
	filePath := folderPath + "/" + fileName + ".go"
	fullPackageName := folderPath
	packageName := extractPackageName(folderPath)

	template := `package ` + packageName + `

import (
	"log/slog"
)

type STRUCT_NAME struct {
	logger *slog.Logger
}

func NewSTRUCT_NAME(
	logger *slog.Logger,
) *STRUCT_NAME {
	return &STRUCT_NAME{
		logger: logger,
	}
}

`
	template = strings.ReplaceAll(template, "STRUCT_NAME", structName)

	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		log.Printf("Error while creating struct: %s", err)
		return err
	}

	if err := registerStructToApp(fullPackageName, packageName, structName); err != nil {
		log.Printf("Error while register struct %s", err)
		return err
	}

	return nil
}
