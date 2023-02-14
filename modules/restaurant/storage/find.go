package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

func (sql *sqlStore) FindRestaurantWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurants, error) {
	var data restaurantmodel.Restaurants

	if err := sql.db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
