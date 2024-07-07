package providers

import (
	"log/slog"
	"os"
)

func ProvideLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
