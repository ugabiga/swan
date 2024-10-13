package app

import (
	"context"
	"github.com/ugabiga/swan/core"
	"log"
)

func NewApp() *core.App {
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

func RunApp() {
	ctx := context.Background()
	newApp := NewApp().Provide()
	if err := newApp.Start(ctx); err != nil {
		log.Fatal(err)
	}
	if err := newApp.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
