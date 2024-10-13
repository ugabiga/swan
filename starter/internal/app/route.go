package app

import (
	"github.com/ugabiga/swan/core"
)

func SetRouteHTTPServer(
	server *core.Server,
) {
	e := server.HTTPServer()
	_ = e.Group("")
}
