package config

import (
	"log/slog"

	"github.com/ugabiga/swan/core"
)

func InvokeToSetCronTabRouter(
	logger *slog.Logger,
	cronTab *core.CronTab,
) {
	cronTab.RegisterCronJob("* * * * *", func() error {
		logger.Info("Cron job 1")
		return nil
	})
}
