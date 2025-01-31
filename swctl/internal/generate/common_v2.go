package generate

import (
	"os"
	"strings"

	"github.com/ugabiga/swan/swctl/internal/utils"
)

const (
	RepoAddr        = "https://github.com/ugabiga/swan.git"
	BoostrapDirName = "bootstrap"
	BootstrapPath   = "swan/bootstrap"
)

func RegisterInvokeInContainer(folderPath, structName, funcName string) error {
	containerPath := ContainerPath
	invokeFunc := "fx.Invoke"

	bytes, err := os.ReadFile(containerPath)
	if err != nil {
		return err
	}
	containerFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	packagePath := moduleName + "/" + folderPath
	if !strings.Contains(containerFileContents, packagePath) {
		containerFileContents = strings.ReplaceAll(containerFileContents, "import (", "import (\n\t\""+packagePath+"\"")
	}

	if strings.Contains(containerFileContents, invokeFunc+"(") {
		if strings.Contains(containerFileContents, structName+"."+funcName) {
			return nil
		}

		containerFileContents = strings.Replace(containerFileContents, invokeFunc+
			"(", invokeFunc+
			"(\n\t\t\t"+structName+"."+funcName+",", 1)

		if err = os.WriteFile(containerPath, []byte(containerFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}

func RegisterStructToContainer(folderPath, packageName, funcName string) error {
	containerPath := ContainerPath
	moduleName := utils.RetrieveModuleName()
	provideFunc := "fx.Provide"

	bytes, err := os.ReadFile(containerPath)
	if err != nil {
		return err
	}
	containerFileContents := string(bytes)

	packagePath := moduleName + "/" + folderPath
	if !strings.Contains(containerFileContents, packagePath) {
		containerFileContents = strings.ReplaceAll(containerFileContents,
			"import (",
			"import (\n\t\""+packagePath+"\"",
		)
	}

	if strings.Contains(containerFileContents, provideFunc+"(") {
		if strings.Contains(containerFileContents, packageName+"."+funcName) {
			return nil
		}

		containerFileContents = strings.Replace(containerFileContents, provideFunc+
			"(", provideFunc+
			"(\n\t\t\t"+packageName+"."+funcName+",", 1)

		if err = os.WriteFile(containerPath, []byte(containerFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}
