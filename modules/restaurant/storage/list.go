package restaurantstorage

import (
	"context"
	common2 "golang_01/common"
	"golang_01/modules/restaurant/model"
)

func (s *sqlStore) ListDataByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common2.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurants, error) {
	offset := (paging.Page - 1) * paging.Limit
	var data []restaurantmodel.Restaurants

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurants{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter.RestaurantId; v > 0 {
		db = db.Where("id = ?", v)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common2.ErrDB(err)
	}

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, common2.ErrDB(err)
	}

	return data, nil
}
