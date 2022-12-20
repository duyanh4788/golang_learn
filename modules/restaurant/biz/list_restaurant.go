package restaurantbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurants, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurants, error) {

	result, err := biz.store.ListDataByConditions(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetListEntity(restaurantmodel.EntityName, err)
	}

	return result, err
}
