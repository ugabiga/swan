package core

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func InvokeSetMainCommand(
	logger *slog.Logger,
	server *Server,
	command *Command,
) {
	command.registerMainCommand(
		&cobra.Command{
			Use:   "main",
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
	crontab *CronTab,
	command *Command,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "cron",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				crontab.Start()
			},
		},
	)
}
