package restaurantbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
	"log"
)

type FindRestaurantStore interface {
	FindRestaurantWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)
}

type findRestaurantBiz struct {
	store          FindRestaurantStore
	restaurantRepo ListRestaurantRepo
}

func NewFindRestaurantBiz(store FindRestaurantStore, restaurantRepo ListRestaurantRepo) *findRestaurantBiz {
	return &findRestaurantBiz{store: store, restaurantRepo: restaurantRepo}
}

func (biz *findRestaurantBiz) FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurants, error) {
	data, err := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrDisableStatus(restaurantmodel.EntityName, data.Name, err)
	}

	if biz.restaurantRepo != nil {
		var restaurant restaurantmodel.Filter
		var paging common.Paging

		restaurant.RestaurantId = id
		paging.Fulfill()

		mapLike, err := biz.restaurantRepo.ListRestaurant(ctx, &restaurant, &paging)

		if err != nil {
			log.Println("Cannot get like count", err)
		}

		if v := mapLike; v != nil {
			for _, item := range mapLike {
				if data.Id == item.Id {
					data.LikeCount = item.LikeCount
				}
			}
		}
	}

	return data, nil
}
