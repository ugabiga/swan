package providers

import (
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/starter/internal/example"
)

func InvokeSetRouteHTTPServer(
	server *core.Server,
	exampleHandler *example.Handler,
) {
	e := server.HTTPServer()

	apiGroup := e.Group("/api")

	exampleHandler.SetRoutes(apiGroup)
}
