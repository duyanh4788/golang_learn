package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/component/asyncjob"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurantlike/model"
	"log"
)

type UserUnLikedRestaurantStore interface {
	FindUserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	UserUnLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type IndAndDeCreaseRestaurant interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikedRestaurantBiz struct {
	store    UserUnLikedRestaurantStore
	inDstore IndAndDeCreaseRestaurant
}

func NewUserUnLikedRestaurantBiz(store UserUnLikedRestaurantStore, inDstore IndAndDeCreaseRestaurant) *userUnLikedRestaurantBiz {
	return &userUnLikedRestaurantBiz{store: store, inDstore: inDstore}
}

func (biz *userUnLikedRestaurantBiz) UserUnLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (string, error) {
	restaurantLike, err := biz.store.FindUserLikeRestaurant(ctx, data)

	if err != nil {
		if err != common.RecordNotFound {
			return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "unliked")
		}
		return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "unliked")
	}

	if restaurantLike.RestaurantId != data.RestaurantId {
		return restaurantlikemodel.NIL, common.NewUnAuthorized(nil, "data not found", "dataNotFound")
	}

	if err := biz.store.UserUnLikedRestaurant(ctx, restaurantLike); err != nil {
		return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "unliked")
	}

	go func() {
		defer common.Recover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			if err := biz.inDstore.DecreaseLikeCount(ctx, data.RestaurantId); err != nil {
				log.Println(err)
				return err
			}
			return nil
		}, asyncjob.WithName("DecreaseLikeCount"))

		if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
			log.Println(err)
		}
	}()

	return restaurantlikemodel.UNLIKE, nil
}
