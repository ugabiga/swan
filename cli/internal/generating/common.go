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

func registerStructToAppInvoker(fullPackageName, packageName, structName string) error {
	appFilePath := "./internal/config/app.go"
	appRegisterProvidersFunc := "app.RegisterInvokers"

	bytes, err := os.ReadFile(appFilePath)
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

		if err = os.WriteFile(appFilePath, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}
