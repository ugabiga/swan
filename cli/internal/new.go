package internal

import (
	"fmt"
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

func CreateNew(
	appName string,
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

	install(appName)

	if err := cleanKeepFiles(appName); err != nil {
		return err
	}

	return nil
}

func install(appName string) {
	commands := []string{
		"go get ./...",
		"go mod tidy",
		"go mod download",
	}

	for _, command := range commands {
		cmd := exec.Command(command)
		cmd.Dir = fmt.Sprintf("./%s", appName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
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
