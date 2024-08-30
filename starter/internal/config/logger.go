package config

import (
	"log/slog"

	"github.com/ugabiga/swan/utl"
)

func ProvideLogger() *slog.Logger {
	return utl.NewCharmLogger(utl.LoggerOption{
		Level: slog.LevelDebug,
	})
}
