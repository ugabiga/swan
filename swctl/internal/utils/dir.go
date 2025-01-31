package utils

import "os"

func IsFolderExists(folderPath string) bool {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func IfFolderNotExistsCreate(folderPath string) error {
	if !IsFolderExists(folderPath) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return err
		}
	}

	return nil
}
