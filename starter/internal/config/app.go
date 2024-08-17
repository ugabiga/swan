package config

import (
	"github.com/ugabiga/swan/core"
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

	//HTTP Server
	app.RegisterProviders(
		func() core.ServerConfig {
			return core.ServerConfig{
				Addr: env.Addr,
			}
		},
		core.NewServer,
	)

	//Default Providers
	app.RegisterProviders(
		ProvideEventEmitter,
		ProvideEventPubSubContainer,
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	//Invoke
	app.RegisterInvokers(
		InvokeSetRouteHTTPServer,
		InvokeSetEventRouter,
		InvokeListenForEvents,
	)

	//Command
	app.RegisterCommands(
		Commands,
	)

	return app
}
