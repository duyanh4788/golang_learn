package restaurantbiz

import (
	"context"
	"golang_01/common"
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
	restaurant, err := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
	}

	if err := biz.store.UpdateRestaurant(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
