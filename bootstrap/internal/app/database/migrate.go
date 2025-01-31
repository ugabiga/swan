package database

import (
	"errors"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/bootstrap/internal/common/dir"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func extractDSN(dialector gorm.Dialector) string {
	switch d := dialector.(type) {
	case *postgres.Dialector:
		return d.DSN
	case *mysql.Dialector:
		return d.DSN
	case *sqlite.Dialector:
		return "sqlite://" + d.DSN
	default:
		return ""
	}
}

func migrateCreate(name string) (bool, error) {
	atlasCmd := exec.Command("atlas", "migrate", "create", "-dir", "-ext", "sql", "-name", name)
	if err := atlasCmd.Run(); err != nil {
		fmt.Println("failed to create migration", err)
		fmt.Println("to install atlas visit https://atlasgo.io/getting-started/")

		return false, err
	}

	createMigrationCmd := exec.Command("atlas", "migrate", "diff", name, "--env", "gorm")
	if err := createMigrationCmd.Run(); err != nil {
		fmt.Println("failed to create migration", err)
		return false, err
	}
	return true, nil
}

func migrateUp(dialector gorm.Dialector) (bool, error) {
	migrationPath := dir.ProjectRoot() + "/migrations"
	dsn := extractDSN(dialector)

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn,
	)
	if err != nil {
		return false, err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func migrateDown(dialector gorm.Dialector) (bool, error) {
	migrationPath := dir.ProjectRoot() + "/migrations"
	dsn := extractDSN(dialector)

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn,
	)
	if err != nil {
		return false, err
	}

	err = m.Down()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func SetCommands(
	db *gorm.DB,
	logger *slog.Logger,
	cmd *cobra.Command,
) {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Database migration commands",
	}

	var migrateCreateCmd = &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new database migration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				logger.Error("migration name is required")
				return
			}

			dialector := db.Dialector
			result, err := migrateCreate(dialector, args[0])
			if err != nil {
				logger.Error("migration creation failed", "err", err)
				return
			}
			if result {
				logger.Info("migration creation completed successfully")
			}
		},
	}

	var migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			dialector := db.Dialector
			result, err := migrateUp(dialector)
			if err != nil {
				logger.Error("migration failed", "err", err)
				return
			}
			if result {
				logger.Info("migration completed successfully")
			} else {
				logger.Info("no migration needed")
			}
		},
	}

	var migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Rollback database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			dialector := db.Dialector
			result, err := migrateDown(dialector)
			if err != nil {
				logger.Error("migration rollback failed", "err", err)
				return
			}
			if result {
				logger.Info("migration rollback completed successfully")
			} else {
				logger.Info("no migration to rollback")
			}
		},
	}

	migrateCmd.AddCommand(migrateCreateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	cmd.AddCommand(migrateCmd)
}
