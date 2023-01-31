package restaurantstorage

import (
	"context"
	"golang_01/common"
	"golang_01/modules/restaurant/model"
)

func (s *sqlStore) ListDataByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurants, error) {
	var data []restaurantmodel.Restaurants

	db := s.db

	db = db.Table(restaurantmodel.Restaurants{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter.RestaurantId; v > 0 {
		db = db.Where("id = ?", v)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if fk := paging.FakeCursor; fk != "" {
		if uid, err := common.FromBase58(fk); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
