package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

func (sql *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := sql.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
