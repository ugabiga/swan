package app

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/ugabiga/swan/bootstrap/internal/app/config"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

const eventLogGroupName = "watermill_event"

func NewEventRouter(logger *slog.Logger) (*message.Router, error) {
	logger = logger.WithGroup(eventLogGroupName)

	router, err := message.NewRouter(
		message.RouterConfig{},
		watermill.NewSlogLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	router.AddMiddleware(middleware.Recoverer)

	return router, err
}

func RunEventRouter(logger *slog.Logger, cfg config.Config, router *message.Router) {
	if !cfg.EventConfig.Enabled {
		return
	}

	go func() {
		ctx := context.Background()
		if err := router.Run(ctx); err != nil {
			logger.Error(err.Error())
		}
	}()
}

func NewEventPublisher(channel *gochannel.GoChannel) (message.Publisher, error) {
	return channel, nil
}

func NewEventSubscriber(channel *gochannel.GoChannel) (message.Subscriber, error) {
	return channel, nil
}

func NewEventChannel(logger *slog.Logger) *gochannel.GoChannel {
	logger = logger.WithGroup(eventLogGroupName)

	return gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewSlogLogger(logger),
	)
}
