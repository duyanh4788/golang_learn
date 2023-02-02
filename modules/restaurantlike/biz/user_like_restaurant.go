package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	FindUserLikeRestaurant(ctx context.Context, id int) (*restaurantlikemodel.RestaurantLike, error)
	UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
	DeleteUserLikeRestaurant(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (string, error) {
	restaurantLike, err := biz.store.FindUserLikeRestaurant(ctx, data.RestaurantId)

	if err != nil {
		if err != common.RecordNotFound {
			if err := biz.store.UserLikeRestaurant(ctx, data); err != nil {
				return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err)
			}
			return restaurantlikemodel.LIKE, nil
		}
		return restaurantlikemodel.NIL, common.ErrIntenval(err)
	}

	if restaurantLike.RestaurantId > 0 {
		if err := biz.store.DeleteUserLikeRestaurant(ctx, restaurantLike.RestaurantId); err != nil {
			return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err)
		}
		return restaurantlikemodel.UNLIKE, nil
	}

	return restaurantlikemodel.NIL, nil
}
