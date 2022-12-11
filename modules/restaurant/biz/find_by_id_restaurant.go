package restaurantbiz

import (
	"context"
	restaurantmodel "golang_01/modules/restaurant/model"
)

type FindRestaurantStore interface {
	FindRestaurantWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)
}

type findRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

func (biz *findRestaurantBiz) FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurants, error) {
	data, err := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}
