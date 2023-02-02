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
	store     FindRestaurantStore
	likeStore LikeStore
}

func NewFindRestaurantBiz(store FindRestaurantStore, likeStore LikeStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store, likeStore: likeStore}
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

	var ids []int

	ids = append(ids, data.Id)

	mapLike, err := biz.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println("Cannot get like count", err)
	}

	if v := mapLike; v != nil {
		data.LikeCount = mapLike[data.Id]
	}

	return data, nil
}
