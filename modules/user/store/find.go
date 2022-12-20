package userstore

import (
	"context"
	"golang_01/common"
	"golang_01/modules/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrCannotFindEntity(usermodel.EntityName, err)
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
