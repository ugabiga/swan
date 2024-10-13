package app

import (
	"github.com/ugabiga/swan/core"
	"log/slog"
)

func SetRouteHTTPServer(
	logger *slog.Logger,
	server *core.Server,
) {
	e := server.HTTPServer()
	g := e.Group("")

	logger.Info("RouteHTTPServer", "group", g)
}
