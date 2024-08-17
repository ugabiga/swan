package main

import (
	"github.com/ugabiga/swan/starter/internal/config"
)

// @title			STARTER_PLACEHOLDER
// @version			0.1.0
// @description		STARTER_PLACEHOLDER
// @host			localhost:8080
// @BasePath
func main() {
	if err := config.ProvideApp().Run(); err != nil {
		panic(err)
	}
}
