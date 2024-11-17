package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func NewServer(logger *slog.Logger) *echo.Echo {
	e := echo.New()

	e.HidePort = true
	e.HideBanner = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(c.Request().Context(), slog.LevelDebug, "Request",
					slog.String("method", c.Request().Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(c.Request().Context(), slog.LevelError, "Request Error",
					slog.String("method", c.Request().Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	return e
}

func SetCommands(
	e *echo.Echo,
	logger *slog.Logger,
	cmd *cobra.Command,
) {
	var serverCmd = &cobra.Command{
		Use:   "server [flags]",
		Short: "Start the HTTP server with optional configuration",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			if port == "" {
				port = "8080"
			}
			addr := ":" + port
			logger.Info("Server Start", "addr", addr)

			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()

			// Start server
			go func() {
				if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("error while running http server", "err", err)
				}
			}()

			// Wait for interrupt signal to gracefully shutdowns the server with a timeout of 10 seconds.
			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := e.Shutdown(ctx); err != nil {
				logger.Error("error while shutdown", "err", err)
			}
			logger.Info("Server Shutdown")
		},
	}

	serverCmd.Flags().String("port", "", "port to run server on")

	var routesCmd = &cobra.Command{
		Use:   "routes",
		Short: "Display all registered HTTP routes and their methods",
		Run: func(cmd *cobra.Command, args []string) {
			for _, route := range e.Routes() {
				logger.Info("", slog.String("method", route.Method), slog.String("path", route.Path))
			}
		},
	}

	serverCmd.AddCommand(routesCmd)

	cmd.AddCommand(serverCmd)
}
