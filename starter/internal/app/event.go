package app

import (
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/core/pubsub"
)

func registerEvents(app *core.App, env *EnvironmentVariables) {
	app.RegisterInvokers(
		SetEventMiddleware,
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
		ProvideEventPubSubContainer,
	)

	if env.EventDriver == "none" {
		return
	}

	registerEvents(app, env)

	if env.EventDriver != "channel" {
		return
	}

	app.RegisterInvokers(
		core.InvokeListenForEvents,
	)

}

func ProvideEventPubSubContainer(env *EnvironmentVariables) (pubsub.Container, error) {
	return pubsub.NewContainer(
		pubsub.ContainerConfig{
			EventDriver: env.EventDriver,
		},
	)
}
