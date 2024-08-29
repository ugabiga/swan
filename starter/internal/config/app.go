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

	app.RegisterInvokers(
		example.InvokeToSetCronTab,
		example.InvokeSetMainCommand,
		example.InvokeSetExampleCommand,
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
		ProvideCronTab,
		ProvideCommand,
		ProvideEventEmitter,
		ProvideEventPubSubContainer,
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	//Invoke
	app.RegisterInvokers(
		InvokeToStartCronTab,
		InvokeSetRouteHTTPServer,
		InvokeSetEventRouter,
		InvokeListenForEvents,
	)

	return app
}
