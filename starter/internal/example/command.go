package example

import "log/slog"

func InvokeCommand(
	logger *slog.Logger,
) {
	logger.Info("Command example")
}
