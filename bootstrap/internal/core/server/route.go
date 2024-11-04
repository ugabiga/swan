package server

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/swan/bootstrap/internal/core/auth"
)

func SetRoutes(
	authManger *auth.Manager,
	openAPIHandler *OpenAPIHandler,
	staticHandler *StaticHandler,
	e *echo.Echo,
) {
	_ = auth.Middleware(authManger)

	rootGroup := e.Group("")
	staticHandler.SetMiddleware(rootGroup)
	openAPIHandler.SetRouters(rootGroup)

	api := e.Group("/api")

	_ = api.Group("/v1")
}
