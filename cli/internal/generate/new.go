package generate

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"path/filepath"
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

	if err := cleanKeepFiles(appName); err != nil {
		return err
	}

	return nil
}
