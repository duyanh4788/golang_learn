package restaurantbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurants, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurants, error) {

	result, err := biz.repo.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
