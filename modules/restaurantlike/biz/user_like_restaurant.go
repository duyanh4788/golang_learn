package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	if err := biz.store.UserLikeRestaurant(ctx, data); err != nil {
		return common.ErrCannotLike(restaurantmodel.EntityName, err)
	}
	return nil
}
