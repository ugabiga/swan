package app

import (
	"testing"

	"github.com/ugabiga/swan/core"
	"go.uber.org/fx"
)

type TestApp struct {
	Deps TestAppDependencies
}

type TestAppDependencies struct {
	fx.In

	//Put your dependencies for the test container here
}

func NewTestApp(deps TestAppDependencies) *TestApp {
	return &TestApp{
		Deps: deps,
	}
}

func ProvideTestApp(t *testing.T) *TestApp {
	core.LoadEnv(".env.test")

	var testContainer *TestApp

	app := NewApp()
	app.RegisterProviders(
		NewTestApp,
	)

	app.Invokers = append(app.Invokers, func(c *TestApp) {
		testContainer = c
	})

	if err := app.Invoke(); err != nil {
		t.Fatal(err)
	}

	return testContainer
}
