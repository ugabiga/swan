package providers

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

	//Event
	app.RegisterProviders(
		ProvidePubSubContainer,
		core.NewEventEmitter,
	)

	app.RegisterProviders(
		ProvideEnvironmentVariables,
		ProvideLogger,
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

	//Invoke
	app.RegisterInvokers(
		InvokeSetRouteHTTPServer,
		InvokeSetEventRouter,
		InvokeListenForEvents,
	)

	//Command
	app.RegisterCommands(
		map[string]any{
			"example": example.NewCommand,
		},
	)

	return app
}
