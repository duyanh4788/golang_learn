package restaurantbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	err := biz.store.CreateRestaurant(ctx, data)

	return err
}
