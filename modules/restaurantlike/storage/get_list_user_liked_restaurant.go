package restaurantlikestorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang_01/common"
	"golang_01/modules/restaurantlike/model"
	"time"
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

func (sql sqlStore) GetUserLikedRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var data []restaurantlikemodel.RestaurantLike

	db := sql.db

	db = db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		db = db.Where("restaurant_id = ?", v.RestaurantId)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if fk := paging.FakeCursor; fk != "" {

		timerCreated, err := time.Parse(common.TimeLayout, string(base58.Decode(fk)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timerCreated.Format("2006-01-02 15:04:05"))

	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	user := make([]common.SimpleUser, len(data))

	for i, item := range data {
		data[i].User.CreatedAt = item.CreatedAt
		data[i].User.UpdatedAt = nil
		user[i] = *data[i].User

		if i == len(data)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(common.TimeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return user, nil
}
