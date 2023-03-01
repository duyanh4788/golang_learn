package subscriber

import (
	"context"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/storage"
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
