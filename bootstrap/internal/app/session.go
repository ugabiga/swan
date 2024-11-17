package app

import (
	"github.com/gorilla/sessions"
	"github.com/ugabiga/swan/bootstrap/internal/app/config"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

func NewSessionStore(cfg config.Config, db *gorm.DB) (sessions.Store, error) {
	secret := cfg.OAuthConfig.SessionSecret
	maxAge := 86400 * 30
	secure := cfg.OAuthConfig.SessionSecure

	switch cfg.OAuthConfig.SessionDriver {
	case "gorm":
		return newGormStore(db, secret, maxAge, secure), nil
	default:
		return newCookieSessionStore(secret, maxAge, secure), nil
	}
}

func newCookieSessionStore(secret string, maxAge int, secure bool) sessions.Store {
	store := sessions.NewCookieStore([]byte(secret))

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
	}

	return store
}

func newGormStore(db *gorm.DB, secret string, maxAge int, secure bool) sessions.Store {
	store := gormstore.NewOptions(db,
		gormstore.Options{
			TableName:       "http_sessions",
			SkipCreateTable: false,
		},
		[]byte(secret),
	)

	store.SessionOpts = &sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
	}

	return store
}
