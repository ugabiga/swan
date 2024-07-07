package providers

import (
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/starter/internal/example"
)

func ProvideAppAndRun() error {
	env := ProvideEnvironmentVariables()
	c := core.NewApp()

	c.RegisterProviders(
		func() core.ServerConfig {
			return core.ServerConfig{
				Addr: env.Addr,
			}
		},
		core.NewServer,
	)

	c.RegisterProviders(
		ProvideEnvironmentVariables,
		ProvideLogger,
	)

	c.RegisterProviders(
		example.NewHandler,
	)

	c.RegisterInvokers(
		InvokeSetRouteHTTPServer,
	)

	return c.Run()
}
