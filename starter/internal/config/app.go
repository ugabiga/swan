package config

import (
	"github.com/ugabiga/swan/core"
)

func ProvideApp() *core.App {
	app := core.NewApp()
	env := ProvideEnvironmentVariables()

	app.RegisterProviders(
	//...
	)

	app.RegisterInvokers(
	//...
	)

	app.RegisterInvokers(
		SetRouteHTTPServer,
	)

	app.RegisterProviders(
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	InitializeConfigs(app, env)
	InitializeCommands(app, env)
	InitializeCron(app, env)
	InitializeEvent(app, env)

	return app
}
