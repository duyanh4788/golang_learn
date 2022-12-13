package restaurantbiz

import (
	"context"
	"golang_01/component/common"
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
		if err != common.RecordNotFound {
			return common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
	}

	if err := biz.store.DeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
