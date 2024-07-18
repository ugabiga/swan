package pubsub

import (
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

type Container interface {
	NewSubscriber() message.Subscriber
	NewPublisher() message.Publisher
}

var (
	ErrRedisConfigIsEmpty = errors.New("redis config is empty")
	ErrSQLConfigIsEmpty   = errors.New("sql config is empty")
)

type ContainerConfig struct {
	EventDriver string  `json:"event_driver"`
	RedisAddr   *string `json:"redis_addr"`
	RedisDB     *int    `json:"redis_db"`
	SQLDBType   *string `json:"sql_db_type"`
	SQLUser     *string `json:"sql_user"`
	SQLPass     *string `json:"sql_pass"`
	SQLHost     *string `json:"sql_addr"`
	SQLPort     *string `json:"sql_port"`
	SQLDBName   *string `json:"sqldb_name"`
}

func NewContainer(
	config ContainerConfig,
) (Container, error) {

	switch config.EventDriver {
	case "channel":
		return NewChannel(), nil
	case "redis":
		if config.RedisAddr == nil || config.RedisDB == nil {
			return nil, ErrRedisConfigIsEmpty
		}

		return NewRedis(redis.NewClient(&redis.Options{
			Addr: *config.RedisAddr,
			DB:   *config.RedisDB,
		})), nil
	case "sql":
		if config.SQLDBType == nil ||
			config.SQLUser == nil ||
			config.SQLPass == nil ||
			config.SQLHost == nil ||
			config.SQLPort == nil ||
			config.SQLDBName == nil {
			return nil, ErrSQLConfigIsEmpty
		}

		return NewSQL(
			*config.SQLDBType,
			createDB(
				*config.SQLDBType,
				*config.SQLUser,
				*config.SQLPass,
				*config.SQLHost,
				*config.SQLPort,
				*config.SQLDBName,
			),
		), nil

	default:
		return NewChannel(), nil
	}
}
