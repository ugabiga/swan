package config

import (
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/core/pubsub"
)

func registerEvents(app *core.App, env *EnvironmentVariables) {
	app.RegisterInvokers(
	//...
	)
}

func SetEventMiddleware(
	eventEmitter *core.EventEmitter,
) {
	eventEmitter.AddMiddleware(
		middleware.Recoverer,
	)
}

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
		SetEventMiddleware,
	)

	registerEvents(app, env)

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
