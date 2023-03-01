package pubsublocal

import (
	"context"
	"golang_01/common"
	"golang_01/pubsub"
	"log"
	"sync"
)

const n = 10000

// A pub sub run locally (in memory)
// It has a queue (buffer channel) at it's core and many group of subscribers
// Because we want to send a message with a specific topic for many subscribers in a group can handle

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, n),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}
	pb.run()

	return pb
}

func (lps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.Recover()
		lps.messageQueue <- data
		log.Println("New event published", data.String(), "data", data.Data())
	}()

	return nil
}

func (lps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	lps.locker.Lock()

	if value, ok := lps.mapChannel[topic]; ok {
		value = append(lps.mapChannel[topic], c)
		lps.mapChannel[topic] = value
	} else {
		lps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	lps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := lps.mapChannel[topic]; ok {
			for i := range lps.mapChannel[topic] {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					lps.locker.Lock()
					lps.mapChannel[topic] = chans
					lps.locker.Unlock()

					break
				}
			}
		}
	}
}

func (lps *localPubSub) run() error {
	log.Println("PubSub Started")

	go func() {
		for {
			mess := <-lps.messageQueue
			log.Println("Message dequeue", mess)

			if subs, ok := lps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
			//else {
			//	lps.messageQueue <- mess
			//}
		}
	}()

	return nil
}
