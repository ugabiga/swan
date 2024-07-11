package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/ugabiga/swan/core"
)

type Channel struct {
	pubSub *gochannel.GoChannel
}

func NewChannel() core.PubSubInstance {
	//TODO: Add slog
	logger := watermill.NewStdLogger(false, false)

	return &Channel{
		pubSub: gochannel.NewGoChannel(
			gochannel.Config{},
			logger,
		),
	}
}

func (pubSub Channel) NewSubscriber() message.Subscriber {
	return pubSub.pubSub
}

func (pubSub Channel) NewPublisher() message.Publisher {
	return pubSub.pubSub
}
