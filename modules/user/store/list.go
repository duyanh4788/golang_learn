package userstore

import (
	"context"
	"golang_01/common"
	"golang_01/modules/admin/model"
	"golang_01/modules/user/model"
)

func (s *sqlStore) ListUsersByAdmin(
	ctx context.Context,
	filter *adminmodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {
	var user []usermodel.User

	db := s.db

	db = db.Table(usermodel.User{}.TableName())

	if v := filter.UserId; v > 0 {
		db = db.Where("id = ?", v)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if fk := paging.FakeCursor; fk != "" {
		if uid, err := common.FromBase58(fk); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return user, nil
}
