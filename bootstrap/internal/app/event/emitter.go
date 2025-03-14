package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventEmitter struct {
	publisher message.Publisher
}

func NewEventEmitter(publisher message.Publisher) *EventEmitter {
	return &EventEmitter{
		publisher: publisher,
	}
}

func (e *EventEmitter) Emit(topic string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return e.EmitRaw(topic, data)
}

func (e *EventEmitter) EmitRaw(topic string, payload []byte) error {
	msg := message.NewMessage(watermill.NewUUID(), payload)
	msg.Metadata.Set("topic", topic)

	return e.publisher.Publish(topic, msg)
}
