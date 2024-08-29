package example

import (
	"log/slog"

	"github.com/ugabiga/swan/core"
)

func InvokeToSetCronTab(
	cronTab *core.CronTab,
	logger *slog.Logger,
) {
	cronTab.RegisterCronJob("* * * * * *", func() error {
		logger.Info("Cron job 1")
		return nil
	})

	cronTab.RegisterCronJob("* * * * * *", func() error {
		logger.Info("Cron job 2")
		return nil
	})
}