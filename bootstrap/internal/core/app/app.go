package app

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/ugabiga/swan/bootstrap/internal/common/dir"
	"gorm.io/gorm"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	app := fx.New(
		fx.NopLogger,
		provide(),
		invoke(),
		entry(),
	)

	return app
}

type TestApp struct {
	Deps TestAppDependencies
}

type TestAppDependencies struct {
	fx.In

	Logger *slog.Logger
	DB     *gorm.DB
}

func NewTestApp() *TestApp {
	if err := godotenv.Load(dir.ProjectRoot() + "/" + ".test.env"); err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}

	var testApp *TestApp

	fx.New(
		fx.NopLogger,
		provide(),
		invoke(),
		fx.Provide(
			func(deps TestAppDependencies) *TestApp {
				return &TestApp{
					Deps: deps,
				}
			},
		),
		fx.Invoke(func(a *TestApp) {
			testApp = a
		}),
	)

	return testApp
}
