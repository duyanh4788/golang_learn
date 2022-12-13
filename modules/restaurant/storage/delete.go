package restaurantstorage

import (
	"context"
	restaurantmodel "golang_01/modules/restaurant/model"
)

func (s *sqlStore) DeleteRestaurant(
	ctx context.Context,
	cond map[string]interface{},
) error {
	if err := s.db.
		Table(restaurantmodel.Restaurants{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return err
	}

	return nil
}
