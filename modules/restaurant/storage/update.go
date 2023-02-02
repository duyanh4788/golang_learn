package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {

	if err := s.db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := *s.db

	if err := db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id =?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := *s.db

	if err := db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id =?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
