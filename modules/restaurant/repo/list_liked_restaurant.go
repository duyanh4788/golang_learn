package restaurantrepo

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

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type ListRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeStore) *ListRestaurantRepo {
	return &ListRestaurantRepo{store: store, likeStore: likeStore}
}

func (repo *ListRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurants, error) {

	result, err := repo.store.ListDataByConditions(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotGetListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	//mapLikes, err := repo.likeStore.GetRestaurantLikes(ctx, ids)
	//
	//if err != nil {
	//	log.Println("Cannot get like count", err)
	//}
	//
	//if v := mapLikes; v != nil {
	//	for i, item := range result {
	//		result[i].LikeCount = mapLikes[item.Id]
	//	}
	//}

	return result, nil
}
