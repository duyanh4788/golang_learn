package restaurantlikestorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

func (sql *sqlStore) UserUnLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	var restaurantLike restaurantlikemodel.RestaurantLike

	if err := sql.db.Where("restaurant_id = ? AND user_id = ?", data.RestaurantId, data.UserId).Delete(&restaurantLike).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
