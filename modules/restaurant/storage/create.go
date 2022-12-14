package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
