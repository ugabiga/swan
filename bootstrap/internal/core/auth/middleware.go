package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const CtxSessionUserKey = "ctx-session-user"

func Middleware(manager *Manager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionUser, err := manager.sessionUser(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			c.Set(CtxSessionUserKey, sessionUser)
			return next(c)
		}
	}
}
