package config

import (
	"log/slog"

	"github.com/ugabiga/swan/core"
)

func InvokeToSetCleanup(
	logger *slog.Logger,
	cleanup *core.Cleanup,
) {
	cleanup.RegisterCleanup(func() error {
		logger.Info("Cleanup 1")
		return nil
	})
}
