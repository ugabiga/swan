package generating

import (
	"errors"
	"fmt"
	"github.com/ugabiga/swan/cli/internal/utils"
	"os"
	"strings"
)

var (
	ErrEntNotInitialized = errors.New("ent not initialized")
)

func InitializeEntORM() error {
	if err := createDBClient(); err != nil {
		return err
	}

	if err := addMigrationSetupToEnv(); err != nil {
		return err
	}

	if err := addMakefileCommands(); err != nil {
		return err
	}

	return nil
}

func addMakefileCommands() error {
	makefilePath := "Makefile"
	migrationCommands := `

migrate-diff:
	@make ent-gen && atlas migrate diff "$(name)" --dir "file://$(MIGRATION_PATH)" --to "ent://internal/ent/schema" --dev-url "$(DEVELOP_DB_URL)"

migrate-up:
	@atlas migrate apply --dir "file://$(MIGRATION_PATH)" --url "$(DB_URL)"

migrate-down:
	@atlas migrate down --dir "file://$(MIGRATION_PATH)" --url "$(DB_URL)" --dev-url "$(DEVELOP_DB_URL)"
`
	file, err := os.OpenFile(makefilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(migrationCommands); err != nil {
		return err
	}

	return nil
}

func addMigrationSetupToEnv() error {
	envFilePath := ".env"
	migrationPath := "migrations"
	dbURL := "postgres://postgres:postgres@localhost:5432/ent?sslmode=disable"
	developDBURL := "docker://postgres/15/dev?search_path=public"

	file, err := os.OpenFile(envFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("\nMIGRATION_PATH=%s\n", migrationPath)); err != nil {
		return err
	}

	if _, err := file.WriteString(fmt.Sprintf("DB_URL=%s\n", dbURL)); err != nil {
		return err
	}

	if _, err := file.WriteString(fmt.Sprintf("DEVELOP_DB_URL=%s\n", developDBURL)); err != nil {
		return err
	}

	return nil
}

func createDBClient() error {
	folderPath := "internal/client"
	structName := "EntClient"
	fileName := "ent"
	filePath := folderPath + "/" + fileName + ".go"
	fullPackageName := folderPath
	packageName := extractPackageName(folderPath)

	template := `package ` + packageName + `

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/url"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	ENT_PACKAGE_PATH

	"modernc.org/sqlite"
)

var (
	ErrUnknownDriver             = errors.New("unknown driver")
	ErrFailedOpeningConnection   = errors.New("failed opening connection")
	ErrFailedToEnableForeignKeys = errors.New("failed to enable foreign keys")
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})

	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		if err := conn.Close(); err != nil {
			return nil, err
		}
		return nil, ErrFailedToEnableForeignKeys
	}

	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

type EntClientConfig struct {
	URL string
}

func NewEntClient(logger *slog.Logger, cfg EntClientConfig) (*ent.Client, error) {
	parts := strings.Split(cfg.URL, "://")
	if len(parts) != 2 {
		log.Fatalf("invalid ent url: %s", cfg.URL)
	}

	driverVal := parts[0]
	onlyURL := parts[1]
	switch driverVal {
	case "sqlite":
		return newSqliteEntClient(logger, onlyURL)
	case "postgres":
		return newPostgresSQLClient(logger, cfg.URL)
	case "mysql":
		return newMySQLClient(logger, cfg.URL)
	default:
		logger.Error("invalid ent driver", slog.String("driver", driverVal))
		return nil, ErrUnknownDriver
	}

	return nil, ErrFailedOpeningConnection
}

func convertMySQLURL(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	userInfo := u.User
	host := u.Hostname()
	port := u.Port()
	dbName := strings.TrimPrefix(u.Path, "/")

	convertedURL := fmt.Sprintf("%s@tcp(%s:%s)/%s", userInfo.String(), host, port, dbName)

	query := u.Query()
	for key, values := range query {
		for _, value := range values {
			convertedURL += fmt.Sprintf("&%s=%s", key, value)
		}
	}

	return convertedURL, nil
}

func newMySQLClient(logger *slog.Logger, url string) (*ent.Client, error) {
	convertedURL, err := convertMySQLURL(url)
	if err != nil {
		logger.Error("failed converting mysql url: %v", err)
		return nil, err
	}

	driver, err := entsql.Open(dialect.MySQL, convertedURL)
	if err != nil {
		logger.Error("failed connecting to mysql: %v", err)
		return nil, err
	}

	client := ent.NewClient(ent.Driver(driver))

	return client, nil
}

func newPostgresSQLClient(logger *slog.Logger, url string) (*ent.Client, error) {
	entDriver, err := entsql.Open(dialect.Postgres, url)
	if err != nil {
		logger.Error("failed connecting to postgres: %v", err)
		return nil, err
	}

	client := ent.NewClient(ent.Driver(entDriver))

	return client, nil
}

func newSqliteEntClient(logger *slog.Logger, url string) (*ent.Client, error) {
	client, err := ent.Open(dialect.SQLite, url)
	if err != nil {
		if err := client.Close(); err != nil {
			return nil, err
		}
		logger.Error("failed opening connection to sqlite: %v", err)
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	return client, nil
}

`
	entClientFolder := "internal/ent"

	// check if the folder exists
	if _, err := os.Stat(entClientFolder); os.IsNotExist(err) {
		return ErrEntNotInitialized
	}

	// get the module name
	moduleName := utils.RetrieveModuleName()
	entClientPackagePath := moduleName + "/" + entClientFolder
	template = strings.ReplaceAll(template, "ENT_PACKAGE_PATH", `"`+entClientPackagePath+`"`)

	//check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			fmt.Printf("Error while creating folder: %s", err)
			return err
		}
	}

	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		fmt.Printf("Error while creating struct: %s", err)
		return err
	}

	if err := registerStructToApp(fullPackageName, packageName, structName); err != nil {
		fmt.Printf("Error while register struct %s", err)
		return err
	}
	return nil
}
