package entdb

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

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

// func NewEntClient(logger *slog.Logger, cfg config.Config) (*ent.Client, error) {
// 	url := cfg.DatabaseConfig.URL
// 	parts := strings.Split(url, "://")
// 	if len(parts) != 2 {
// 		log.Fatalf("invalid ent url: %s", url)
// 	}

// 	driverVal := parts[0]
// 	onlyURL := parts[1]
// 	switch driverVal {
// 	case "sqlite":
// 		return newSqliteEntClient(logger, onlyURL)
// 	case "postgres":
// 		return newPostgresSQLClient(logger, url)
// 	case "mysql":
// 		return newMySQLClient(logger, url)
// 	default:
// 		logger.Error("invalid ent driver", slog.String("driver", driverVal))
// 		return nil, ErrUnknownDriver
// 	}
// }

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

// func newMySQLClient(logger *slog.Logger, url string) (*ent.Client, error) {
// 	convertedURL, err := convertMySQLURL(url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	driver, err := entsql.Open(dialect.MySQL, convertedURL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := ent.NewClient(ent.Driver(driver))

// 	return client, nil
// }

// func newPostgresSQLClient(logger *slog.Logger, url string) (*ent.Client, error) {
// 	entDriver, err := entsql.Open(dialect.Postgres, url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := ent.NewClient(ent.Driver(entDriver))

// 	return client, nil
// }

// func newSqliteEntClient(logger *slog.Logger, url string) (*ent.Client, error) {
// 	client, err := ent.Open(dialect.SQLite, url)
// 	if err != nil {
// 		if err := client.Close(); err != nil {
// 			return nil, err
// 		}
// 		return nil, err
// 	}

// 	return client, nil
// }
