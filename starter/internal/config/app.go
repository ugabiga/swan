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
		ProvideEventPubSubContainer,
		ProvideEnvironmentVariables,
		ProvideLogger,
		core.NewEventEmitter,
		core.NewCronTab,
		core.NewCommand,
	)

	//Invoke
	app.RegisterInvokers(
		InvokeSetRouteHTTPServer,
		InvokeSetEventRouter,
		InvokeToSetCronTabRouter,
		core.InvokeListenForEvents,
		core.InvokeSetCronCommand,
		core.InvokeSetMainCommand,
	)

	return app
}
