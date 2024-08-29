package config

import (
	"fmt"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/core/pubsub"
)

func ProvideEventPubSubContainer(env *EnvironmentVariables) (pubsub.Container, error) {
	return pubsub.NewContainer(
		pubsub.ContainerConfig{
			EventDriver: env.EventDriver,
			RedisAddr:   &env.EventRedisAddr,
			RedisDB:     &env.EventRedisDB,
			SQLDBType:   &env.EventSQLDBType,
			SQLUser:     &env.EventSQLUser,
			SQLPass:     &env.EventSQLPass,
			SQLHost:     &env.EventSQLHost,
			SQLPort:     &env.EventSQLPort,
			SQLDBName:   &env.EventSQLDBName,
		},
	)
}

func InvokeSetEventRouter(
	logger *slog.Logger,
	eventEmitter *core.EventEmitter,
) {
	eventEmitter.AddOneWayHandler(
		"exampleHandler",
		"example",
		func(msg *message.Message) error {
			logger.Info("Received message",
				slog.Any("message", msg),
				slog.String("payload", fmt.Sprintf("%s", msg.Payload)),
			)

			return nil
		},
	)
}
