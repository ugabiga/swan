package main

import "github.com/ugabiga/swan/starter/internal/providers"

//	@title			Starter
//	@version		0.1.0
//	@description	Swan
//	@termsOfService

//	@contact.name	API Support
//	@contact.url
//	@contact.email

//	@license.name
//	@license.url

// @host		localhost:8080
// @BasePath	/
func main() {
	if err := providers.ProvideAppAndRun(); err != nil {
		panic(err)
	}
}
