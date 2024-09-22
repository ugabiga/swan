package config

import (
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/starter/internal/example"
)

func ProvideApp() *core.App {
	app := core.NewApp()
	env := ProvideEnvironmentVariables()

	app.RegisterProviders(
		example.NewHandler,
		example.NewService,
	)

	app.RegisterInvokers(
		example.InvokeSetExampleCommand,
	)

	InitializeEvent(app, env)
	InitializeCore(app, env)
	InitializeConfigs(app, env)

	return app
}
