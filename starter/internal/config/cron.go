package config

import (
	"log/slog"

	"github.com/ugabiga/swan/core"
)

func ProvideCronTab(logger *slog.Logger) *core.CronTab {
	return core.NewCronTab(logger)
}

func InvokeToStartCronTab(logger *slog.Logger, cronTab *core.CronTab) {
	if err := cronTab.Start(); err != nil {
		logger.Error("Error", slog.Any("error", err))
		return
	}
}
