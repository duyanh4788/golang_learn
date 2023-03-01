package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurantlike/model"
	"golang_01/pubsub"
)

type UserLikedRestaurantStore interface {
	FindUserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	UserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

//type InAndDeCreaseRestaurant interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikedRestaurantBiz struct {
	store UserLikedRestaurantStore
	//inDstore InAndDeCreaseRestaurant
	pubSub pubsub.Pubsub
}

func NewUserLikedRestaurantBiz(
	store UserLikedRestaurantStore,
	//inDstore InAndDeCreaseRestaurant,
	pubSub pubsub.Pubsub,
) *userLikedRestaurantBiz {
	return &userLikedRestaurantBiz{
		store: store,
		//inDstore: inDstore,
		pubSub: pubSub,
	}
}

func (biz *userLikedRestaurantBiz) UserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (string, error) {
	restaurantLike, err := biz.store.FindUserLikedRestaurant(ctx, data)

	if restaurantLike != nil {
		return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "liked")
	}

	if err != nil {
		if err != common.RecordNotFound {
			if err := biz.store.UserLikedRestaurant(ctx, data); err != nil {
				return restaurantlikemodel.NIL, common.ErrCannotLike(restaurantmodel.EntityName, err, "liked")
			}

			biz.pubSub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))
			//go func() {
			//	defer common.Recover()
			//	job := asyncjob.NewJob(func(ctx context.Context) error {
			//		if err := biz.inDstore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			//			log.Println(err)
			//			return err
			//		}
			//		return nil
			//	}, asyncjob.WithName("IncreaseLikeCount"))
			//
			//	if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
			//		log.Println(err)
			//	}
			//}()
			return restaurantlikemodel.LIKE, nil
		}
		return restaurantlikemodel.NIL, common.ErrIntenval(err)
	}

	return restaurantlikemodel.NIL, nil
}
