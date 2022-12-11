package restaurantbiz

import (
	"context"
	"golang_01/modules/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindRestaurantWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)

	UpdateRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	_, err := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.UpdateRestaurant(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
