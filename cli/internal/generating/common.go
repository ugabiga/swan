package generating

import (
	"github.com/ugabiga/swan/cli/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

func extractPackageName(path string) string {
	packageName := filepath.Base(path)
	packageName = strings.TrimSuffix(packageName, "/")

	return packageName
}

// RegisterStructToAppInvoker registers a struct to the app invoker
// config/app path : internal/config/app.go
func registerStructToAppInvoker(filePath, fullPackageName, packageName, structName string) error {
	appRegisterProvidersFunc := "app.RegisterInvokers"

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	// Format the package path
	var packagePath string
	packagePath = moduleName + "/" + fullPackageName
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, appRegisterProvidersFunc+"(") {
		appFileContents = strings.Replace(appFileContents, appRegisterProvidersFunc+
			"(", appRegisterProvidersFunc+
			"(\n\t\t"+packageName+"."+structName+",", 1)

		if err = os.WriteFile(filePath, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}
