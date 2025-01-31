package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func RetrieveModuleName() string {
	// Read the go.mod file
	bytes, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	contents := string(bytes)

	// Retrieve the module name
	lines := strings.Split(contents, "\n")
	moduleName := strings.Split(lines[0], " ")[1]

	return moduleName
}

func ExtractPackageName(path string) string {
	packageName := filepath.Base(path)
	packageName = strings.TrimSuffix(packageName, "/")

	return packageName
}
