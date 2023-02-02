package restaurantlikestorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

func (sql *sqlStore) FindUserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error) {

	var restaurantLike restaurantlikemodel.RestaurantLike

	if err := sql.db.Where("restaurant_id = ? AND user_id = ?", data.RestaurantId, data.UserId).First(&restaurantLike).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	return &restaurantLike, nil
}

func (sql *sqlStore) UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	if err := sql.db.
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlStore) UserUnLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	var restaurantLike restaurantlikemodel.RestaurantLike

	if err := sql.db.Where("restaurant_id = ? AND user_id = ?", data.RestaurantId, data.UserId).Delete(&restaurantLike).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
