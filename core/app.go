package core

import (
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
		useDependencyLogger: false,
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
		InvokeSetWorkerCommand,
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

func (c *App) Run() {
	var command *Command
	var logger *slog.Logger

	options := []fx.Option{
		fx.Provide(c.Providers...),
		fx.Invoke(c.Invokers...),
		fx.Invoke(func(
			invokedLogger *slog.Logger,
			invokedCommand *Command,
		) {
			logger = invokedLogger
			command = invokedCommand
		}),
	}

	if !c.useDependencyLogger {
		options = append(options, fx.NopLogger)
	}

	fx.New(options...)

	if err := command.Run(); err != nil {
		logger.Error("Failed to run command", slog.Any("error", err))
	}
}

func (c *App) cleanUp(logger *slog.Logger, cleanup *Cleanup) {
	logger.Debug("Running cleanup...")
	if cleanup != nil {
		if err := cleanup.Run(); err != nil {
			logger.Error("Failed to run cleanup", slog.Any("error", err))
		}
	}
}
