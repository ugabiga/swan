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

	// Events
	app.RegisterProviders(
		core.NewEventEmitter,
	)

	app.RegisterInvokers(
		core.InvokeListenForEvents,
	)

	// Core
	app.RegisterProviders(
		ProvideEventPubSubContainer,
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	app.RegisterInvokers(
		InvokeToSetCleanup,
		InvokeSetRouteHTTPServer,
		InvokeSetEventRouter,
		InvokeToSetCronTabRouter,
	)

	app.SetUseDependencyLogger(false)

	ProvideConfigs(app, env)

	return app
}
