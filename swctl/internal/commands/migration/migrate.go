package migration

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func extractDSN() string {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get DATABASE_URL from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not found in environment")
	}

	return dsn
}

func migratePath() string {
	return "migrations"
}

func migrateDrop() (bool, error) {
	migrationPath := migratePath()
	dsn := extractDSN()

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn,
	)
	if err != nil {
		return false, err
	}

	err = m.Drop()
	if err != nil {
		return false, err
	}

	return true, nil
}

func migrateUp() (bool, error) {
	migrationPath := migratePath()
	dsn := extractDSN()

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn,
	)

	if err != nil {
		return false, err
	}

	err = m.Up()
	if err != nil {
		return false, err
	}

	return true, nil
}

func migrateDown() (bool, error) {
	migrationPath := migratePath()
	dsn := extractDSN()

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn,
	)

	if err != nil {
		return false, err
	}

	err = m.Down()
	if err != nil {
		return false, err
	}

	return true, nil
}

func migrateForce(version int) (bool, error) {
	migrationPath := migratePath()
	dsn := extractDSN()

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn,
	)

	if err != nil {
		return false, err
	}

	err = m.Force(version)
	if err != nil {
		return false, err
	}

	return true, nil
}
