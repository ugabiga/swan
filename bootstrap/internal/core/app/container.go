package app

import (
	"github.com/ugabiga/swan/bootstrap/internal/core/auth"
	"github.com/ugabiga/swan/bootstrap/internal/core/config"
	"github.com/ugabiga/swan/bootstrap/internal/core/database"
	"github.com/ugabiga/swan/bootstrap/internal/core/server"
	"go.uber.org/fx"
)

func provide() fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewConfig,
			database.NewGormClient,
			auth.NewManager,
			server.NewServer,
			server.NewOpenAPIHandler,
			server.NewStaticHandler,
		),
		fx.Provide(
			NewEventRouter,
			NewEventChannel,
			NewEventPublisher,
			NewEventSubscriber,
		),
		fx.Provide(
			NewSessionStore,
			NewLogger,
			NewCommand,
		),
	)
}

func invoke() fx.Option {
	return fx.Options(
		fx.Invoke(
			database.SetCommands,
		),
		fx.Invoke(
			server.SetRoutes,
			server.SetCommands,
		),
	)
}

func entry() fx.Option {
	return fx.Options(
		//fx.Invoke(RunEventRouter),
		fx.Invoke(RunCommand),
	)
}
