package example

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/core"
)

func InvokeSetMainCommand(
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

func InvokeSetExampleCommand(
	command *core.Command,
	logger *slog.Logger,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "example",
			Short: "example command",
			Run: func(cmd *cobra.Command, args []string) {
				logger.Info("Example command",
					slog.Any("args", args),
				)
			},
		},
	)
}
