package userstore

import (
	"context"
	"golang_01/common"
	"golang_01/modules/user/model"
)

func (s *sqlStore) UpdateStatusUser(ctx context.Context, userStatus *usermodel.UserUpdateStatus) error {
	if err := s.db.Table(usermodel.User{}.TableName()).Where("id = ?", userStatus.UserId).Update("status", userStatus.Status).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
