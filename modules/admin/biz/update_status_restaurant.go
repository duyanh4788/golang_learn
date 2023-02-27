package adminbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

type UpdateStatusRestaurantStore interface {
	UpdateStatusRestaurant(ctx context.Context, restaurant *restaurantmodel.UpdateStatusRestaurant) error
}

type updateStatusRestaurantBiz struct {
	store UpdateStatusRestaurantStore
}

func NewUpdateStatusRestaurantBiz(store UpdateStatusRestaurantStore) *updateStatusRestaurantBiz {
	return &updateStatusRestaurantBiz{store: store}
}

func (biz *updateStatusRestaurantBiz) UpdateRestaurantByAdmin(ctx context.Context, restaurant *restaurantmodel.UpdateStatusRestaurant) error {
	if err := biz.store.UpdateStatusRestaurant(ctx, restaurant); err != nil {
		return common.ErrCannotUpdateEntity(common.CurrentRestaurant, err)
	}

	return nil
}
