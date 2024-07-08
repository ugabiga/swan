package providers

import (
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/starter/internal/example"
)

func ProvideAppAndRun() error {
	env := ProvideEnvironmentVariables()
	app := core.NewApp()

	app.RegisterProviders(
		example.NewHandler,
	)

	app.RegisterProviders(
		func() core.ServerConfig {
			return core.ServerConfig{
				Addr: env.Addr,
			}
		},
		core.NewServer,
	)

	app.RegisterProviders(
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	app.RegisterInvokers(
		InvokeSetRouteHTTPServer,
	)

	return app.Run()
}
