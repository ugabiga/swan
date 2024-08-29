package core

import (
	"log/slog"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type App struct {
	useDependencyLogger bool
	Providers           []any
	Invokers            []any
}

type AppConfig struct {
	Addr string
}

func NewApp() *App {
	a := &App{
		useDependencyLogger: true,
	}

	a.RegisterProviders(
		NewCronTab,
		NewCommand,
		NewServer,
	)

	a.RegisterInvokers(
		InvokeSetCronCommand,
		InvokeSetMainCommand,
	)

	return a
}

func (c *App) SetUseDependencyLogger(useDependencyLogger bool) {
	c.useDependencyLogger = useDependencyLogger
}

func (c *App) RegisterProviders(providers ...any) {
	c.Providers = append(c.Providers, providers...)
}

func (c *App) RegisterInvokers(invokers ...any) {
	c.Invokers = append(c.Invokers, invokers...)
}

func (c *App) Invoke() error {
	fx.New(
		fx.Provide(c.Providers...),
		fx.Invoke(c.Invokers...),
	)

	return nil
}

func (c *App) Run() error {
	options := []fx.Option{
		fx.Provide(c.Providers...),
		fx.Invoke(c.Invokers...),
		fx.Invoke(func(
			logger *slog.Logger,
			command *Command,
			server *Server,
		) {
			if err := command.Run(); err != nil {
				logger.Error("Failed to run command", slog.Any("error", err))
			}
		}),
	}

	if !c.useDependencyLogger {
		options = append(options, fx.NopLogger)
	}

	fx.New(options...)

	return nil
}

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
