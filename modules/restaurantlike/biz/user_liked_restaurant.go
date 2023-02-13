package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/component/asyncjob"
	restaurantmodel "golang_01/modules/restaurant/model"
	"golang_01/modules/restaurantlike/model"
)

type UserLikedRestaurantStore interface {
	FindUserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	UserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type FinAndDeCreaseRestaurant interface {
	IncreaseLikeCount(ctx context.Context, id int) error
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userLikedRestaurantBiz struct {
	store    UserLikedRestaurantStore
	inDstore FinAndDeCreaseRestaurant
}

func NewUserLikedRestaurantBiz(store UserLikedRestaurantStore, inDstore FinAndDeCreaseRestaurant) *userLikedRestaurantBiz {
	return &userLikedRestaurantBiz{store: store, inDstore: inDstore}
}

func (biz *userLikedRestaurantBiz) UserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (string, error) {
	restaurantLike, err := biz.store.FindUserLikeRestaurant(ctx, data)

	if restaurantLike != nil {
		return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "liked")
	}

	if err != nil {
		if err != common.RecordNotFound {
			if err := biz.store.UserLikedRestaurant(ctx, data); err != nil {
				return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "liked")
			}
			go func() {
				recover()
				job := asyncjob.NewJob(func(ctx context.Context) error {
					return biz.inDstore.IncreaseLikeCount(ctx, data.RestaurantId)
				})

				_ = asyncjob.NewGroup(true, job).Run(ctx)
			}()
			return restaurantlikemodel.LIKE, nil
		}
		return restaurantlikemodel.NIL, common.ErrIntenval(err)
	}

	return restaurantlikemodel.NIL, nil
}
