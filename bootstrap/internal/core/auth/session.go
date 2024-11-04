package auth

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/ugabiga/swan/bootstrap/internal/core/config"
)

const (
	SessionKey            = "auth-session"
	SessionKeyUserKey     = "auth-session-user"
	SessionCtxProviderKey = "provider"
)

var (
	ErrUnknownLoginInUser   = errors.New("error unknown login in user")
	ErrSessionValueNotFound = errors.New("session value not found")
)

type SessionUser struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Manager struct {
	sessionStore sessions.Store
}

func NewManager(cfg config.Config, sessionStore sessions.Store) *Manager {
	m := &Manager{sessionStore: sessionStore}
	m.initGoth(sessionStore, cfg)

	return m
}

func (m Manager) initGoth(store sessions.Store, cfg config.Config) {
	//goth use gob to decode session user
	gob.Register(SessionUser{})

	gothic.Store = store
	goth.UseProviders(
		google.New(
			cfg.OAuthConfig.GoogleClientKey,
			cfg.OAuthConfig.GoogleClientSecret,
			cfg.OAuthConfig.GoogleCallbackURL,
			"email", "profile",
		),
	)
}

func (m Manager) CheckOrInitiateLogin(c echo.Context, provider string) error {
	_, err := m.sessionUser(c)
	if err == nil {
		return nil
	}

	if err = m.login(c, provider); err != nil {
		return err
	}

	return nil
}

func (m Manager) sessionUser(c echo.Context) (*SessionUser, error) {
	value, err := m.getSessionValue(c,
		SessionKey,
		SessionKeyUserKey,
	)
	if err != nil {
		return nil, err
	}

	user, ok := value.(SessionUser)
	if ok {
		return &user, nil
	}

	return nil, ErrUnknownLoginInUser
}

func (m Manager) login(c echo.Context, provider string) error {
	m.setCtx(c, SessionCtxProviderKey, provider)
	gothic.BeginAuthHandler(c.Response(), c.Request())

	return nil
}

func (m Manager) Logout(c echo.Context) error {
	if err := m.deleteSessionValue(c, SessionKey, SessionKeyUserKey); err != nil {
		return err
	}

	if err := gothic.Logout(c.Response(), c.Request()); err != nil {
		return err
	}

	return nil
}

func (m Manager) Callback(c echo.Context) (goth.User, error) {
	return gothic.CompleteUserAuth(c.Response(), c.Request())
}

func (m Manager) SetSessionUser(c echo.Context, sessionUser SessionUser) error {
	if err := m.setSessionValue(c, SessionKey, SessionKeyUserKey, sessionUser); err != nil {
		return err
	}

	return nil
}

func (m Manager) setCtx(c echo.Context, key, value any) {
	c.SetRequest(
		c.Request().WithContext(
			context.WithValue(
				c.Request().Context(),
				key,
				value,
			),
		),
	)
}

func (m Manager) setSessionValue(c echo.Context, sessionName, key string, value any) error {
	sess, _ := m.sessionStore.Get(c.Request(), sessionName)
	sess.Values[key] = value
	return sess.Save(c.Request(), c.Response())
}

func (m Manager) getSessionValue(c echo.Context, sessionName, key string) (any, error) {
	sess, err := m.sessionStore.Get(c.Request(), sessionName)
	if err != nil {
		return nil, err
	}

	value, ok := sess.Values[key]
	if !ok {
		return nil, ErrSessionValueNotFound
	}

	return value, nil
}

func (m Manager) deleteSessionValue(c echo.Context, sessionName, key string) error {
	sess, err := m.sessionStore.Get(c.Request(), sessionName)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	sess.Values = make(map[interface{}]interface{})

	return sess.Save(c.Request(), c.Response())
}

func GetUserFromContext(c echo.Context) (*SessionUser, error) {
	// Retrieve the user from the context
	user, ok := c.Get(CtxSessionUserKey).(*SessionUser)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "user not found in context")
	}
	return user, nil
}
