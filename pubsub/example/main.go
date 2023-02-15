package main

import (
	"context"
	"golang_01/pubsub"
	"golang_01/pubsub/pubsublocal"
	"log"
	"time"
)

func main() {
	var localPb pubsub.Pubsub = pubsublocal.NewPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, _ := localPb.Subscribe(context.Background(), topic)
	sub2, _ := localPb.Subscribe(context.Background(), topic)

	localPb.Publish(context.Background(), topic, pubsub.NewMessage("sub1"))
	localPb.Publish(context.Background(), topic, pubsub.NewMessage("sub2"))

	go func() {
		for {
			log.Println("Sending sub1", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Sending sub2", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	//time.Sleep(time.Second * 3)
	//close1()
	//close2()

	localPb.Publish(context.Background(), topic, pubsub.NewMessage("sub3"))

	time.Sleep(time.Second * 3)
}
