package config

import (
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/core/pubsub"
)

func InitializeEvent(app *core.App, env *EnvironmentVariables) {
	app.RegisterProviders(
		core.NewEventEmitter,
	)

	app.RegisterProviders(
		ProvideEventPubSubContainer,
	)

	if env.EventDriver == "none" {
		return
	}

	// The order of invokers is important
	app.RegisterInvokers(
		InvokeSetEventMiddleware,
		InvokeSetEventRouter,
	)

	if env.EventDriver != "channel" {
		return
	}

	// The order of invokers is important
	app.RegisterInvokers(
		core.InvokeListenForEvents,
	)

}

func ProvideEventPubSubContainer(env *EnvironmentVariables) (pubsub.Container, error) {
	return pubsub.NewContainer(
		pubsub.ContainerConfig{
			EventDriver: env.EventDriver,
			RedisAddr:   &env.EventRedisAddr,
			RedisDB:     &env.EventRedisDB,
			SQLDBType:   &env.EventSQLDBType,
			SQLUser:     &env.EventSQLUser,
			SQLPass:     &env.EventSQLPass,
			SQLHost:     &env.EventSQLHost,
			SQLPort:     &env.EventSQLPort,
			SQLDBName:   &env.EventSQLDBName,
		},
	)
}

func InvokeSetEventMiddleware(
	eventEmitter *core.EventEmitter,
) {
	eventEmitter.AddMiddleware(
		middleware.Recoverer,
	)
}

func InvokeSetEventRouter(
	logger *slog.Logger,
	eventEmitter *core.EventEmitter,
) {
	eventEmitter.AddOneWayHandler(
		"exampleHandler",
		"example",
		func(msg *message.Message) error {
			logger.Info("Received message",
				slog.Any("uuid", msg.UUID),
				slog.String("payload", string(msg.Payload)),
			)

			return nil
		},
	)
}
