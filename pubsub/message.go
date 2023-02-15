package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	chanel    Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (msg *Message) String() string {
	return fmt.Sprintf("Message %s", msg.chanel)
}

func (msg *Message) Channel() Topic {
	return msg.chanel
}

func (msg *Message) SetChannel(channel Topic) {
	msg.chanel = channel
}

func (msg *Message) Data() interface{} {
	return msg.data
}
