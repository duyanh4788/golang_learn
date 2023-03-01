package subscriber

import (
	"context"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/storage"
	"golang_01/pubsub"
)

func DecreaseUserUnlikedRestaurant(appContext component.AppContext, ctx context.Context) {
	c, _ := appContext.GetPubSub().Subscribe(ctx, common.TopicUserUnLikeRestaurant)
	store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())

	go func() {
		msg := <-c
		unlikeData := msg.Data().(HasRestaurantId)

		_ = store.DecreaseLikeCount(ctx, unlikeData.GetRestaurantId())
	}()
}

func RunDecreaseUserUnlikedRestaurant(appContext component.AppContext) consumerJob {
	return consumerJob{
		Title: "User unliked Restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
