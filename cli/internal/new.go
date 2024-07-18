package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

const (
	StarterRepo    = "https://github.com/ugabiga/swan.git"
	StarterDirName = "starter"
	StarterPath    = "swan/starter"
)

var (
	ErrorNoYarn = errors.New("yarn not found")
)

func CreateNew(
	appName string,
	addWebProject bool,
) error {
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      StarterRepo,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	if err := removeOtherFiles(appName); err != nil {
		return err
	}

	srcDir := appName + "/" + StarterDirName
	destDir := appName

	files, err := os.ReadDir(srcDir)
	if err != nil {
		fmt.Println("Error reading source directory:", err)
		return nil
	}

	for _, file := range files {
		srcFilePath := filepath.Join(srcDir, file.Name())
		destFilePath := filepath.Join(destDir, file.Name())

		if err := os.Rename(srcFilePath, destFilePath); err != nil {
			fmt.Println("Error moving file:", err)
			return nil
		}
	}

	if err = os.RemoveAll(fmt.Sprintf("./%s/"+StarterDirName, appName)); err != nil {
		return err
	}

	if err := renameGoModuleInGoModFile(appName); err != nil {
		return err
	}

	if err := renameGoModuleInGoFiles(appName); err != nil {
		return err
	}

	if err := setupEnvFile(appName); err != nil {
		return err
	}

	if err := setupMainFile(appName); err != nil {
		return err
	}

	if err := setupDependencies(appName); err != nil {
		return err
	}

	// TODO : Refactor this
	if addWebProject {
		if err := setupDependenciesForWeb(appName); err != nil {
			log.Printf("Error setting up web project dependencies: %v", err)
			log.Printf("Skipping web project dependencies setup")
		}
	} else {
		if err := os.RemoveAll(fmt.Sprintf("./%s/web", appName)); err != nil {
			return err
		}
	}

	if err := cleanKeepFiles(appName); err != nil {
		return err
	}

	return nil
}

func setupEnvFile(appName string) error {
	envExampleFile, err := os.ReadFile("./" + appName + "/.env.example")
	if err != nil {
		return err
	}

	envFileContents := string(envExampleFile)
	envFileContents = strings.ReplaceAll(envFileContents, "starter", appName)

	if err := os.WriteFile("./"+appName+"/.env", []byte(envFileContents), 0644); err != nil {
		return err
	}

	if err := os.WriteFile("./"+appName+"/.env.test", []byte(envFileContents), 0644); err != nil {
		return err
	}

	return nil
}

func setupMainFile(appName string) error {
	filePath := fmt.Sprintf("./%s/cmd/swan/main.go", appName)
	mainFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	mailFileContents := string(mainFile)
	mailFileContents = strings.ReplaceAll(mailFileContents, "STARTER_PLACEHOLDER", appName)

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
		fmt.Sprintf("./%s/cli", appName),
		fmt.Sprintf("./%s/core", appName),
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
	goModContents = strings.ReplaceAll(goModContents, StarterPath, appName)

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

			newContents := strings.ReplaceAll(string(read), StarterPath, appName)

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
