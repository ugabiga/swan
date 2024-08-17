package config

import "github.com/ugabiga/swan/core"

type EnvironmentVariables struct {
	AppName string `validate:"required" env:"APP_NAME" json:"app_name,omitempty"`
	Addr    string `validate:"required" env:"ADDR" json:"addr,omitempty"`

	EventDriver    string `validate:"required" env:"EVENT_DRIVER" json:"event_driver,omitempty"`
	EventRedisAddr string `env:"EVENT_REDIS_ADDR" json:"event_redis_addr,omitempty"`
	EventRedisDB   int    `env:"EVENT_REDIS_DB" json:"event_redis_db,omitempty"`
	EventSQLDBType string `env:"EVENT_SQL_DB_TYPE" json:"event_sql_db_type,omitempty"`
	EventSQLUser   string `env:"EVENT_SQL_USER" json:"event_sql_user,omitempty"`
	EventSQLPass   string `env:"EVENT_SQL_PASS" json:"event_sql_pass,omitempty"`
	EventSQLHost   string `env:"EVENT_SQL_HOST" json:"event_sql_host,omitempty"`
	EventSQLPort   string `env:"EVENT_SQL_PORT" json:"event_sql_port,omitempty"`
	EventSQLDBName string `env:"EVENT_SQL_DB_NAME" json:"event_sql_db_name,omitempty"`
}

func ProvideEnvironmentVariables() *EnvironmentVariables {
	return core.ValidateEnvironmentVariables[EnvironmentVariables]()
}
