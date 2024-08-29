package core

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

type App struct {
	useDependencyLogger bool
	Providers           []any
	Invokers            []any
	cleanUpFuncList     []func()
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

func (c *App) RegisterCleanUp(cleanUps ...func()) {
	c.cleanUpFuncList = append(c.cleanUpFuncList, cleanUps...)
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
			// Defer the cleanup function
			defer func() {
				c.cleanUp()
			}()

			// Create a channel to receive OS signals
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

			// Run a goroutine to wait for the signal
			go func() {
				sig := <-sigChan
				logger.Info("Received signal", slog.Any("signal", sig))
				c.cleanUp()
				os.Exit(0)
			}()

			// Run the command
			if err := command.Run(); err != nil {
				logger.Error("Failed to run command", slog.Any("error", err))
			}
		}),
	}

	if !c.useDependencyLogger {
		options = append(options, fx.NopLogger)
	}

	fx.New(options...).Run()

	return nil
}

func (c *App) cleanUp() {
	for _, cleanUp := range c.cleanUpFuncList {
		cleanUp()
	}
}
