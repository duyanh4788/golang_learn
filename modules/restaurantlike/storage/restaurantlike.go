package restaurantlikestore

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
)

func (sql sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlLike struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}

	var listLike []sqlLike

	if err := sql.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id , count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}
