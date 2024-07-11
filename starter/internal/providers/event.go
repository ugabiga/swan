package providers

import (
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ugabiga/swan/core"
)

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
			)

			return nil
		},
	)
}

func InvokeListenForEvents(
	eventEmitter *core.EventEmitter,
) {
	eventEmitter.ListenForEvents()
}
