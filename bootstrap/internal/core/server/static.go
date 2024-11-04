package server

import (
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ugabiga/swan/bootstrap/web"
)

type StaticHandler struct {
	logger *slog.Logger
}

func NewStaticHandler(logger *slog.Logger) (*StaticHandler, error) {
	return &StaticHandler{
		logger: logger,
	}, nil
}

func (h StaticHandler) SetMiddleware(router *echo.Group) {
	dist := web.GetDistFS()

	setHttpDirMiddleware := func() {
		router.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			HTML5:      true,
			Filesystem: http.Dir("web/dist"),
			Skipper: func(c echo.Context) bool {
				return len(c.Request().URL.Path) >= 4 && c.Request().URL.Path[:4] == "/api"
			},
		}))
	}

	if dist == nil {
		setHttpDirMiddleware()
		return
	}

	distFS, err := fs.Sub(*dist, "dist")
	if err != nil {
		h.logger.Error("set static http.Dir", "err", err)
		setHttpDirMiddleware()
		return
	}

	router.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Filesystem: http.FS(distFS),
		Skipper: func(c echo.Context) bool {
			return len(c.Request().URL.Path) >= 4 && c.Request().URL.Path[:4] == "/api"
		},
	}))
}
