package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
	"gorm.io/gorm"
)

func (sql *sqlStore) UpdateRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {

	if err := sql.db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (sql *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := *sql.db

	if err := db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id =?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (sql *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := *sql.db

	if err := db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id =?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
