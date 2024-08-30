package core

import (
	"context"
	"log"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ugabiga/swan/core/pubsub"
)

func InvokeListenForEvents(
	eventEmitter *EventEmitter,
) {
	go eventEmitter.ListenForEvents()
}

type EventEmitter struct {
	logger          *slog.Logger
	router          *message.Router
	pubSubContainer pubsub.Container
	subscriber      message.Subscriber
	publisher       message.Publisher
}

func NewEventEmitter(
	logger *slog.Logger,
	pubSubContainer pubsub.Container,
) *EventEmitter {
	router, err := message.NewRouter(
		message.RouterConfig{},
		watermill.NewSlogLogger(logger),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &EventEmitter{
		logger:          logger,
		router:          router,
		pubSubContainer: pubSubContainer,
		subscriber:      pubSubContainer.NewSubscriber(),
		publisher:       pubSubContainer.NewPublisher(),
	}
}

func (emitter *EventEmitter) AddOneWayHandler(handlerName, topic string, handler message.NoPublishHandlerFunc) {
	emitter.router.AddNoPublisherHandler(handlerName, topic, emitter.subscriber, handler)
}

func (emitter *EventEmitter) AddTwoWayHandler(handlerName, topicSubscription, topicPublisher string, handler message.HandlerFunc) {
	emitter.router.AddHandler(handlerName, topicSubscription, emitter.subscriber, topicPublisher, emitter.publisher, handler)
}

func (emitter *EventEmitter) AddCustomOneWayHandler(
	handlerName string,
	subscribeTopic string,
	subscriber message.Subscriber,
	handlerFunc message.NoPublishHandlerFunc,
) {
	emitter.router.AddNoPublisherHandler(handlerName, subscribeTopic, subscriber, handlerFunc)
}

func (emitter *EventEmitter) AddCustomTwoWayHandler(
	handlerName string,
	subscribeTopic string,
	subscriber message.Subscriber,
	publishTopic string,
	publisher message.Publisher,
	handlerFunc message.HandlerFunc,
) {
	emitter.router.AddHandler(handlerName, subscribeTopic, subscriber, publishTopic, publisher, handlerFunc)
}

func (emitter *EventEmitter) Emit(topic string, payload []byte) error {
	msg := message.NewMessage(watermill.NewUUID(), payload)
	if err := emitter.publisher.Publish(topic, msg); err != nil {
		return err
	}

	return nil
}

func (emitter *EventEmitter) Run() {
	emitter.logger.Info("Listening for events")

	err := emitter.router.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (emitter *EventEmitter) ListenForEvents() {
	emitter.Run()
}
