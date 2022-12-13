package restaurantstorage

import (
	"context"
	"golang_01/component/common"
	"golang_01/modules/restaurant/model"
)

func (s *sqlStore) FindRestaurantWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurants, error) {
	var data restaurantmodel.Restaurants

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
