package migration

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "migrate:create [name]",
	Short: "Create a new database migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("migration name is required")
			return
		}

		result, err := migrateCreate(args[0])
		if err != nil {
			fmt.Println("migration creation failed", "err", err)
			return
		}
		if result {
			fmt.Println("migration creation completed successfully")
		}
	},
}

var HashCmd = &cobra.Command{
	Use:   "migrate:hash",
	Short: "generate hash for the migration",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := migrateHash()
		if err != nil {
			fmt.Println("migration hash generation failed", "err", err)
			return
		}
		if result {
			fmt.Println("migration hash generation completed successfully")
		}
	},
}

var DropCmd = &cobra.Command{
	Use:   "migrate:drop",
	Short: "drop the migration",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := migrateDrop()
		if err != nil {
			fmt.Println("migration drop failed", "err", err)
			return
		}
		if result {
			fmt.Println("migration hash generation completed successfully")
		}
	},
}

var UpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "up the migration",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := migrateUp()
		if err != nil {
			fmt.Println("migration up failed", "err", err)
			return
		}
		if result {
			fmt.Println("migration up completed successfully")
		}
	},
}

var DownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "down the migration",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := migrateDown()
		if err != nil {
			fmt.Println("migration down failed", "err", err)
			return
		}
		if result {
			fmt.Println("migration down completed successfully")
		}
	},
}

var ForceCmd = &cobra.Command{
	Use:   "migrate:force [version]",
	Short: "force the migration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("migration version is required")
			return
		}

		version, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("migration version is invalid", "err", err)
			return
		}

		result, err := migrateForce(version)
		if err != nil {
			fmt.Println("migration force failed", "err", err)
			return
		}
		if result {
			fmt.Println("migration force completed successfully")
		}
	},
}
