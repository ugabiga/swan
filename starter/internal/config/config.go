package config

import "github.com/ugabiga/swan/core"

func InitializeConfigs(app *core.App, env *EnvironmentVariables) {
	app.RegisterProviders(
		func() core.ServerConfig { return core.ServerConfig{Addr: env.Addr} },
	)
}
