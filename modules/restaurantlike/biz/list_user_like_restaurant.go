package restaurantlikebiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUserLikeRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]common.SimpleUser, error)
}

type ListUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *ListUserLikeRestaurantBiz {
	return &ListUserLikeRestaurantBiz{store: store}
}

func (biz *ListUserLikeRestaurantBiz) ListUserLikeRestaurant(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUserLikeRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
