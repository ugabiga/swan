package providers

import (
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/starter/internal/example"
)

func InvokeSetRouteHTTPServer(
	exampleHandler *example.Handler,
	server *core.Server,
) {
	e := server.HTTPServer()

	group := e.Group("")

	exampleHandler.SetRoutes(group)
}
