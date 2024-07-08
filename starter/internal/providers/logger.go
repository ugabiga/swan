package providers

import (
	"log/slog"

	"github.com/ugabiga/swan/utl"
)

func ProvideLogger() *slog.Logger {
	return utl.NewDefaultLogger(
		&slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		},
	)
}
