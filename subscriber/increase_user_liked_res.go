package subscriber

import (
	"context"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/storage"
	"golang_01/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

func IncreaseUserLikedRestaurant(appContext component.AppContext, ctx context.Context) {
	c, _ := appContext.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())

	go func() {
		defer common.Recover()
		for {
			msg := <-c

			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}

func RunIncreaseUserLikedRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "User liked count Restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnect())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
