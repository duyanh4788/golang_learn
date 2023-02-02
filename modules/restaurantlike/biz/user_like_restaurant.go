package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/component/asyncjob"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	FindUserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
	UserUnLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type InAndDeCreaseRestaurant interface {
	IncreaseLikeCount(ctx context.Context, id int) error
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	inDstore InAndDeCreaseRestaurant
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, inDstore InAndDeCreaseRestaurant) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, inDstore: inDstore}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (string, error) {
	restaurantLike, err := biz.store.FindUserLikeRestaurant(ctx, data)

	if err != nil {
		if err != common.RecordNotFound {
			if err := biz.store.UserLikeRestaurant(ctx, data); err != nil {
				return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err)
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

	if restaurantLike.RestaurantId > 0 {
		if err := biz.store.UserUnLikeRestaurant(ctx, restaurantLike); err != nil {
			return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err)
		}

		go func() {
			recover()
			job := asyncjob.NewJob(func(ctx context.Context) error {
				return biz.inDstore.DecreaseLikeCount(ctx, data.RestaurantId)
			})

			_ = asyncjob.NewGroup(true, job).Run(ctx)
		}()

		return restaurantlikemodel.UNLIKE, nil
	}

	return restaurantlikemodel.NIL, nil
}
