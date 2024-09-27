package core

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerConfig struct {
	Addr string
}

type Server struct {
	e            *echo.Echo
	serverConfig ServerConfig
	logger       *slog.Logger
}

func NewServer(
	serverConfig ServerConfig,
	logger *slog.Logger,
) *Server {
	server := &Server{
		e:            echo.New(),
		serverConfig: serverConfig,
		logger:       logger,
	}

	server.initHTTPServer()

	return server
}

func (s *Server) HTTPServer() *echo.Echo {
	return s.e
}

func (s *Server) StartHTTPServer() error {
	s.PrintAllRoutes()
	s.logger.Info("Starting HTTP server", slog.String("addr", s.serverConfig.Addr))

	return s.e.Start(s.serverConfig.Addr)
}

func (s *Server) PrintAllRoutes() {
	routes := s.e.Routes()

	s.logger.Info("Registered routes")
	for _, route := range routes {
		s.logger.Info("", slog.String("method", route.Method), slog.String("path", route.Path))
	}
}

func (s *Server) Shutdown() error {
	return s.e.Shutdown(context.Background())
}

func (s *Server) initHTTPServer() {
	logger := s.logger

	s.e.HidePort = true
	s.e.HideBanner = true

	s.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelDebug, "Request",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "Request Error",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
}
