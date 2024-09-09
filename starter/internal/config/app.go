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

	// Events order between invokers and providers matter
	app.RegisterInvokers(
		InvokeSetEventMiddleware,
		InvokeSetEventRouter,
		core.InvokeListenForEvents,
	)

	app.RegisterProviders(
		core.NewEventEmitter,
	)

	// Core
	app.RegisterInvokers(
		InvokeToSetCleanup,
		InvokeSetRouteHTTPServer,
		InvokeSetCronTabRouter,
	)

	app.RegisterProviders(
		ProvideEventPubSubContainer,
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	app.SetUseDependencyLogger(false)

	ProvideConfigs(app, env)

	return app
}
