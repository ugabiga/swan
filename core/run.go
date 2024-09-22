package core

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func InvokeSetServerCommand(
	logger *slog.Logger,
	server *Server,
	command *Command,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "server",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				if err := server.StartHTTPServer(); err != nil {
					logger.Error("Failed to start the server", slog.Any("error", err))
				}
			},
		},
	)
}

func InvokeSetCronCommand(
	logger *slog.Logger,
	crontab *CronTab,
	command *Command,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "cron",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				logger.Info("Starting cron")
				crontab.Start()
			},
		},
	)
}

func InvokeSetWorkerCommand(
	logger *slog.Logger,
	eventEmitter *EventEmitter,
	command *Command,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "worker",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				logger.Info("Starting worker")
				eventEmitter.Run()
			},
		},
	)
}
