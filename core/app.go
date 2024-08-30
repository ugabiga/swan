package core

import (
	"context"
	"log/slog"

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
		NewCleanup,
	)

	a.RegisterInvokers(
		InvokeSetServerCommand,
		InvokeSetCronCommand,
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
			lc fx.Lifecycle,
			logger *slog.Logger,
			cleanup *Cleanup,
			command *Command,
			server *Server,
		) {
			lc.Append(fx.Hook{
				OnStart: func(context.Context) error {
					return command.Run()
				},
				OnStop: func(context.Context) error {
					c.cleanUp(logger, cleanup)
					return nil
				},
			})
		}),
	}

	if !c.useDependencyLogger {
		options = append(options, fx.NopLogger)
	}

	fx.New(options...).Run()

	return nil
}

func (c *App) cleanUp(logger *slog.Logger, cleanup *Cleanup) {
	logger.Debug("Running cleanup...")
	if cleanup != nil {
		if err := cleanup.Run(); err != nil {
			logger.Error("Failed to run cleanup", slog.Any("error", err))
		}
	}
}
