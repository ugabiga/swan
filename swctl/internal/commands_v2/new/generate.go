package new

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

const (
	RepoAddr        = "https://github.com/ugabiga/swan.git"
	BoostrapDirName = "bootstrap"
	BootstrapPath   = "swan/bootstrap"
)

func Generate(name string) error {
	if err := cloneRepo(name); err != nil {
		return err
	}

	if err := removeExceptBootstrap(name); err != nil {
		return err
	}

	if err := moveBootstrapFiles(name); err != nil {
		return err
	}

	if err := renameGoModuleInGoModFile(name); err != nil {
		return err
	}

	if err := renameGoModuleInGoFiles(name); err != nil {
		return err
	}

	if err := setupEnvFile(name); err != nil {
		return err
	}

	if err := setupMainFile(name); err != nil {
		return err
	}

	if err := setupDependencies(name); err != nil {
		return err
	}

	if err := cleanKeepFiles(name); err != nil {
		return err
	}

	if err := runPnpmInstall(name); err != nil {
		return err
	}

	return nil
}

func cloneRepo(name string) error {
	_, err := git.PlainClone("./"+name, false, &git.CloneOptions{
		URL: RepoAddr,
	})
	if err != nil {
		return err
	}

	return nil
}

func removeExceptBootstrap(appName string) error {
	fileAndDirNames := []string{
		fmt.Sprintf("./%s/.git", appName),
		fmt.Sprintf("./%s/starter", appName),
		fmt.Sprintf("./%s/swctl", appName),
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

func moveBootstrapFiles(appName string) error {
	srcDir := appName + "/" + BoostrapDirName
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

	if err = os.RemoveAll(fmt.Sprintf("./%s/"+BoostrapDirName, appName)); err != nil {
		return err
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

func runPnpmInstall(appName string) error {
	webDir := filepath.Join(appName, "web")
	os.Chdir(webDir)

	if _, err := exec.Command("pnpm", "install").Output(); err != nil {
		return fmt.Errorf("failed to run pnpm install: %w", err)
	}

	return nil
}
