package providers

import (
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/core/pubsub"
	"github.com/ugabiga/swan/starter/internal/example"
)

func ProvideApp() *core.App {
	env := ProvideEnvironmentVariables()
	app := core.NewApp()

	//Domain
	app.RegisterProviders(
		example.NewHandler,
		example.NewService,
	)

	//Event
	app.RegisterProviders(
		pubsub.NewChannel,
		core.NewEventEmitter,
	)

	app.RegisterProviders(
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	//HTTP Server
	app.RegisterProviders(
		func() core.ServerConfig {
			return core.ServerConfig{
				Addr: env.Addr,
			}
		},
		core.NewServer,
	)

	//Invoke
	app.RegisterInvokers(
		InvokeSetRouteHTTPServer,
		InvokeSetEventRouter,
		InvokeListenForEvents,
	)

	return app
}
