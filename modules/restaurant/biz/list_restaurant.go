package restaurantbiz

import (
	"context"
	common2 "golang_01/common"
	"golang_01/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common2.Paging,
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
	paging *common2.Paging) ([]restaurantmodel.Restaurants, error) {

	result, err := biz.store.ListDataByConditions(ctx, nil, filter, paging)

	if err != nil {
		return nil, common2.ErrCannotGetListEntity(restaurantmodel.EntityName, err)
	}

	return result, err
}
