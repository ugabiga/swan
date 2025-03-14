package app

import (
	"github.com/ugabiga/swan/bootstrap/internal/app/auth"
	"github.com/ugabiga/swan/bootstrap/internal/app/config"
	"github.com/ugabiga/swan/bootstrap/internal/app/database"
	"github.com/ugabiga/swan/bootstrap/internal/app/event"
	"github.com/ugabiga/swan/bootstrap/internal/app/server"
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
			event.NewEventEmitter,
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
			server.SetRouter,
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
