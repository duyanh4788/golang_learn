package restaurantbiz

import (
	"context"
	"errors"
	"golang_01/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindRestaurantWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)

	DeleteRestaurant(
		ctx context.Context,
		cond map[string]interface{},
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	restaurant, err := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if restaurant.Status == 0 {
		return errors.New("restaurant not found")
	}

	if err := biz.store.DeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
