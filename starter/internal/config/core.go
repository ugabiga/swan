package config

import "github.com/ugabiga/swan/core"

func InitializeCore(app *core.App, env *EnvironmentVariables) {
	app.RegisterInvokers(
		InvokeToSetCleanup,
		InvokeSetRouteHTTPServer,
		InvokeSetCronTabRouter,
	)

	app.RegisterProviders(
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

}
