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
		Delete(nil).
		Error; err != nil {
		return err
	}

	return nil
}
