package subscriber

import (
	"context"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/storage"
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
