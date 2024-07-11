package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
	logger watermill.LoggerAdapter
}

func NewRedis(
	redisClient *redis.Client,
) Container {
	logger := watermill.NewStdLogger(false, false)

	return &Redis{
		logger: logger,
		client: redisClient,
	}
}

func (pubSub Redis) NewSubscriber() message.Subscriber {
	subscriber, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client:        pubSub.client,
			Unmarshaller:  redisstream.DefaultMarshallerUnmarshaller{},
			ConsumerGroup: "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	return subscriber
}

func (pubSub Redis) NewPublisher() message.Publisher {
	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client:     pubSub.client,
			Marshaller: redisstream.DefaultMarshallerUnmarshaller{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	return publisher
}
