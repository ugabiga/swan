package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DatabaseConfig DatabaseConfig `json:"database_config"`
	OAuthConfig    OAuthConfig    `json:"oauth_config"`
	EventConfig    EventConfig    `json:"event_config"`
}

type DatabaseConfig struct {
	URL string `json:"url"`
}

type OAuthConfig struct {
	SessionDriver      string `json:"session_driver"`
	SessionSecret      string `json:"session_secret"`
	SessionSecure      bool   `json:"session_secure"`
	GoogleClientKey    string `json:"google_client_key"`
	GoogleClientSecret string `json:"google_client_secret"`
	GoogleCallbackURL  string `json:"google_callback_url"`
}

type EventConfig struct {
	Enabled bool `json:"enabled"`
}

func NewConfig() Config {
	return Config{
		DatabaseConfig: DatabaseConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		OAuthConfig: OAuthConfig{
			SessionDriver: os.Getenv("OAUTH_SESSION_DRIVER"),
			SessionSecret: os.Getenv("OAUTH_SESSION_SECRET"),
			SessionSecure: func() bool {
				if val, err := strconv.ParseBool(os.Getenv("OAUTH_SESSION_SECURE")); err == nil {
					return val
				}
				return true
			}(),
			GoogleClientKey:    os.Getenv("GOOGLE_CLIENT_KEY"),
			GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			GoogleCallbackURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
		},
		EventConfig: EventConfig{
			Enabled: func() bool {
				if val, err := strconv.ParseBool(os.Getenv("EVENT_ENABLED")); err == nil {
					return val
				}
				return false
			}(),
		},
	}
}
