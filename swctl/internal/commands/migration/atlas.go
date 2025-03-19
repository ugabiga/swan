package migration

import (
	"fmt"
	"os/exec"
)

func migrateCreate(name string) (bool, error) {
	atlasCmd := exec.Command("which", "atlas")
	if err := atlasCmd.Run(); err != nil {
		fmt.Println("failed to create migration", err)
		fmt.Println("to install atlas visit https://atlasgo.io/getting-started/")
		return false, err
	}

	createMigrationCmd := exec.Command("atlas", "migrate", "diff", name, "--env", "local")
	output, err := createMigrationCmd.CombinedOutput()
	if err != nil {
		fmt.Println("failed to create migration:", err)
		fmt.Println("command output:", string(output))
		return false, err
	}
	return true, nil
}

func migrateHash() (bool, error) {
	atlasCmd := exec.Command("which", "atlas")
	if err := atlasCmd.Run(); err != nil {
		fmt.Println("failed to create migration", err)
		fmt.Println("to install atlas visit https://atlasgo.io/getting-started/")
		return false, err
	}

	createMigrationCmd := exec.Command("atlas", "migrate", "hash", "--env", "local")
	output, err := createMigrationCmd.CombinedOutput()
	if err != nil {
		fmt.Println("failed to create migration:", err)
		fmt.Println("command output:", string(output))
		return false, err
	}
	return true, nil
}
