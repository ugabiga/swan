package gormdb

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/ugabiga/swan/bootstrap/internal/app/config"
	"github.com/ugabiga/swan/bootstrap/internal/common/dir"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormClient(logger *slog.Logger, cfg config.Config) (*gorm.DB, error) {
	dialector, err := newDialector(cfg)
	if err != nil {
		logger.Error("error while init dialector", "err", err)
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: slogGorm.New(
			slogGorm.WithHandler(logger.Handler()),
			slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelError),
		),
	})
	if err != nil {
		logger.Error("failed to open database connection", "err", err)
		return nil, err
	}

	return db, nil
}

func newDialector(cfg config.Config) (gorm.Dialector, error) {
	dnsParts := strings.Split(cfg.DatabaseConfig.URL, "://")
	if len(dnsParts) != 2 {
		return nil, errors.New("invalid database URL")
	}
	driver := dnsParts[0]
	databaseURL := dnsParts[1]

	switch driver {
	case "sqlite":
		dbPath := dir.ProjectRoot() + "/" + databaseURL
		return sqlite.Open(dbPath), nil
	default:
		return nil, errors.New("invalid database driver")
	}
}
