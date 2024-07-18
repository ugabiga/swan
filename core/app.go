package core

import (
	"fmt"
	"log/slog"

	"go.uber.org/fx"
)

type App struct {
	Providers []any
	Invokers  []any
	Commands  map[string]any
}

type AppConfig struct {
	Addr string
}

func NewApp() *App {
	return &App{}
}

func (c *App) RegisterProviders(providers ...any) {
	c.Providers = append(c.Providers, providers...)
}

func (c *App) RegisterInvokers(invokers ...any) {
	c.Invokers = append(c.Invokers, invokers...)
}

func (c *App) RegisterCommands(commands map[string]any) {
	c.Commands = commands
}

func (c *App) Run() error {
	fx.New(
		fx.Provide(c.Providers...),
		fx.Invoke(c.Invokers...),
		fx.Invoke(func(
			lc fx.Lifecycle,
			logger *slog.Logger,
			server *Server,
		) {
			if err := server.StartHTTPServer(); err != nil {
				logger.Error("Failed to start the server", slog.Any("error", err))
			}
		}),
	)

	return nil
}

func (c *App) Invoke() error {
	fx.New(
		fx.Provide(c.Providers...),
		fx.Invoke(c.Invokers...),
	)

	return nil
}

func (c *App) RunCommand(commandStr string) error {
	command := c.Commands[commandStr]
	if command == nil {
		return fmt.Errorf("command %s not found", commandStr)
	}

	fx.New(
		fx.Provide(c.Providers...),
		fx.Invoke(c.Invokers...),
		fx.Invoke(command),
	)

	return nil
}
