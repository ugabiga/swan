package example

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/core"
)

func InvokeNewCommand(
	command *core.Command,
	logger *slog.Logger,
) {
	command.RegisterMainCommand(
		&cobra.Command{
			Use:   "main",
			Short: "Main command",
			Run: func(cmd *cobra.Command, args []string) {
				logger.Info("Command example")
			},
		},
	)
}
