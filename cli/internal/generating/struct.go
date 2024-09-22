package generating

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateStruct(path, name string) {
	folderPath := "internal/" + path

	//Check if folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			panic(err)
		}
	}

	if err := createStruct(folderPath, name); err != nil {
		panic(err)
	}
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

func extractPackageName(path string) string {
	packageName := filepath.Base(path)
	packageName = strings.TrimSuffix(packageName, "/")

	return packageName
}
