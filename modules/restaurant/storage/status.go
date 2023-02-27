package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

func (sql *sqlStore) UpdateStatusRestaurant(ctx context.Context, restaurant *restaurantmodel.UpdateStatusRestaurant) error {
	if err := sql.db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id =?", restaurant.RestaurantId).
		Update("status", restaurant.Status).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
