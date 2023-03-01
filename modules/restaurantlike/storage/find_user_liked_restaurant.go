package restaurantlikestorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

func (sql *sqlStore) FindUserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error) {

	var restaurantLike restaurantlikemodel.RestaurantLike

	if err := sql.db.Where("restaurant_id = ? AND user_id = ?", data.RestaurantId, data.UserId).First(&restaurantLike).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	return &restaurantLike, nil
}
