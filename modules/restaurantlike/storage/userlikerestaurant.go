package restaurantlikestorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

func (sql *sqlStore) UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	var restaurantLike restaurantlikemodel.RestaurantLike

	if err := sql.db.Where("restaurant_id = ? AND user_id = ?", data.RestaurantId, data.UserId).First(&restaurantLike).Error; err != nil {
		if err != common.RecordNotFound {
			if err := sql.db.
				Create(data).Error; err != nil {
				return common.ErrDB(err)
			}
			return nil
		}
		return common.ErrDB(err)
	}

	if restaurantLike.RestaurantId >= 0 {
		if err := sql.db.Where("restaurant_id = ?", data.RestaurantId).Delete(&restaurantLike).Error; err != nil {
			return common.ErrDB(err)
		}
		return nil
	}
	return nil
}
