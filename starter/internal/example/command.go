package example

import "log/slog"

func NewCommand(
	logger *slog.Logger,
) {
	logger.Info("Command example")
}
