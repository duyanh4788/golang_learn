package restaurantstorage

import (
	"context"
	restaurantmodel "golang_01/modules/restaurant/model"
)

func (s *sqlStore) UpdateRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {

	if err := s.db.Where(conditions).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
