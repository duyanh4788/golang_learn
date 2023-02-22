package userstore

import (
	"context"
	"golang_01/common"
	"golang_01/modules/user/model"
)

func (s *sqlStore) UpdateProfile(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdate,
) error {

	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil

}

func (s *sqlStore) UpdatePassWord(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdatePassWord,
) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
