package restaurantlikestorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

func (sql *sqlStore) UserLikedRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	if err := sql.db.
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
