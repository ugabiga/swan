package database

import (
	"errors"
	"log/slog"

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

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	cmd.AddCommand(migrateCmd)
}
