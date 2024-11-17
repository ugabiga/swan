package app

import (
	"log/slog"
	"os"

	charmlogger "github.com/charmbracelet/log"
)

func NewLogger() *slog.Logger {
	level := slog.LevelDebug

	charmLoggerStyle := charmlogger.DefaultStyles()

	charmLogger := charmlogger.NewWithOptions(
		os.Stdout,
		charmlogger.Options{
			Level:           charmlogger.Level(level),
			ReportTimestamp: true,
			ReportCaller:    true,
		},
	)
	charmLogger.SetStyles(charmLoggerStyle)

	return slog.New(charmLogger)
}
