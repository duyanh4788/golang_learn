package restaurantbiz

import (
	"context"
	common2 "golang_01/common"
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
		if err != common2.RecordNotFound {
			return nil, common2.ErrCannotFindEntity(restaurantmodel.EntityName, err)
		}
		return nil, common2.ErrCannotFindEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common2.ErrCannotFindEntity(restaurantmodel.EntityName, nil)
	}

	return data, nil
}
