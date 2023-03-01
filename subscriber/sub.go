package subscriber

import (
	"context"
	"golang_01/pubsub"
)

type consumerJob struct {
	Title string
	Hdl   func(ctx context.Context, message *pubsub.Message) error
}
