package generate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ugabiga/swan/swctl/internal/utils"
)

func extractPackageName(path string) string {
	packageName := filepath.Base(path)
	packageName = strings.TrimSuffix(packageName, "/")

	return packageName
}

func registerToInvoker(invokerFilePath, fullPackageName, packageName, structName string) error {
	appRegisterProvidersFunc := "app.RegisterInvokers"

	bytes, err := os.ReadFile(invokerFilePath)
	if err != nil {
		return err
	}

	invokerFileContent := string(bytes)
	targetModuleName := utils.RetrieveModuleName()

	// Format the package path
	var packagePath string
	packagePath = targetModuleName + "/" + fullPackageName
	invokerFileContent = strings.ReplaceAll(invokerFileContent, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(invokerFileContent, appRegisterProvidersFunc+"(") {
		invokerFileContent = strings.Replace(
			invokerFileContent,
			appRegisterProvidersFunc+"(", appRegisterProvidersFunc+"(\n\t\t"+packageName+"."+structName+",",
			1,
		)

		if err = os.WriteFile(invokerFilePath, []byte(invokerFileContent), 0644); err != nil {
			return err
		}
	}

	return nil
}

func registerStructToApp(fullPackageName, packageName, structName string) error {
	appFilePath := AppPath
	appRegisterProvidersFunc := "app.RegisterProviders"

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
			"(\n\t\t"+packageName+".New"+structName+",", 1)

		if err = os.WriteFile(appFilePath, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}

func registerHandlerToApp(folderPath, domainName string) error {
	appFilePath := ContainerPath
	appRegisterProvidersFunc := "fx.Provide"

	bytes, err := os.ReadFile(appFilePath)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	packagePath := moduleName + "/" + folderPath
	log.Printf("packagePath: %s", packagePath)
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, appRegisterProvidersFunc+"(") {
		appFileContents = strings.Replace(appFileContents, appRegisterProvidersFunc+
			"(", appRegisterProvidersFunc+
			"(\n\t\t"+domainName+".NewHandler,", 1)

		if err = os.WriteFile(appFilePath, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}

func registerHandlerToRoute(folderPath, domainName string) error {
	routerFile := RouterPath
	routerInvokeFunc := "SetRouter"

	bytes, err := os.ReadFile(routerFile)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	// Format the package path
	packagePath := moduleName + "/" + folderPath
	log.Printf("packagePath: %s", packagePath)
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, routerInvokeFunc+"(") {
		appFileContents = strings.Replace(appFileContents,
			routerInvokeFunc+"(",
			routerInvokeFunc+"(\n\t"+domainName+"Handler *"+domainName+".Handler,",
			1,
		)

		if err = os.WriteFile(routerFile, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	if strings.Contains(appFileContents, "}") {
		appFileContents = strings.Replace(appFileContents,
			"}",
			"\t"+domainName+"Handler.SetRoutes(api)\n}",
			1,
		)
		if err = os.WriteFile(routerFile, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}

func setupEnvFile(appName string) error {
	envExampleFile, err := os.ReadFile("./" + appName + "/.example.env")
	if err != nil {
		return err
	}

	envFileContents := string(envExampleFile)
	envFileContents = strings.ReplaceAll(envFileContents, "bootstrap", appName)

	if err := os.WriteFile("./"+appName+"/.env", []byte(envFileContents), 0644); err != nil {
		return err
	}

	if err := os.WriteFile("./"+appName+"/.test.env", []byte(envFileContents), 0644); err != nil {
		return err
	}

	return nil
}

func setupMainFile(appName string) error {
	filePath := fmt.Sprintf("./%s/cmd/app/main.go", appName)
	mainFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	mailFileContents := string(mainFile)
	mailFileContents = strings.ReplaceAll(mailFileContents, "BOOTSTRAP_PLACEHOLDER", appName)

	if err := os.WriteFile(filePath, []byte(mailFileContents), 0644); err != nil {
		return err
	}

	return nil
}

func setupDependencies(appName string) error {
	commands := []*exec.Cmd{
		exec.Command("go", "get", "./..."),
		exec.Command("go", "mod", "tidy"),
		exec.Command("go", "mod", "download"),
	}

	for _, command := range commands {
		command.Dir = "./" + appName
		if _, err := command.Output(); err != nil {
			return err
		}
	}

	return nil
}

func setupDependenciesForWeb(appName string) error {
	commands := []*exec.Cmd{
		exec.Command("yarn"),
	}

	_, err := exec.LookPath("yarn")
	if err != nil {
		return ErrorNoYarn
	}

	for _, command := range commands {
		command.Dir = "./" + appName + "/web"
		if _, err := command.Output(); err != nil {
			return err
		}
	}

	return nil
}

func removeOtherFiles(appName string) error {
	fileAndDirNames := []string{
		fmt.Sprintf("./%s/.git", appName),
		fmt.Sprintf("./%s/starter", appName),
		fmt.Sprintf("./%s/cli", appName),
		fmt.Sprintf("./%s/core", appName),
		fmt.Sprintf("./%s/utl", appName),
		fmt.Sprintf("./%s/go.work.sum", appName),
		fmt.Sprintf("./%s/LICENSE", appName),
		fmt.Sprintf("./%s/README.md", appName),
	}

	for _, fileOrDir := range fileAndDirNames {
		if err := os.RemoveAll(fileOrDir); err != nil {
			return err
		}
	}

	return nil
}

func renameGoModuleInGoModFile(appName string) error {
	goMod := fmt.Sprintf("./%s/go.mod", appName)
	goModFile, err := os.ReadFile(goMod)
	if err != nil {
		panic(err)
	}

	goModContents := string(goModFile)
	goModContents = strings.ReplaceAll(goModContents, BootstrapPath, appName)

	if err = os.WriteFile(goMod, []byte(goModContents), 0644); err != nil {
		return err
	}

	return nil
}

func renameGoModuleInGoFiles(appName string) error {
	return filepath.Walk("./"+appName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		matched := strings.HasSuffix(path, ".go")

		if matched {
			read, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			newContents := strings.ReplaceAll(string(read), BootstrapPath, appName)

			err = os.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func cleanKeepFiles(appName string) error {
	err := filepath.Walk("./"+appName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".keep") {
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
