package userbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/user/model"
)

type UpdateProfileStore interface {
	UpdateProfile(ctx context.Context, id int, data *usermodel.UserUpdate) error
}

type updateProfileBiz struct {
	store UpdateProfileStore
}

func NewUpdateProfileBiz(store UpdateProfileStore) *updateProfileBiz {
	return &updateProfileBiz{store: store}
}

func (biz *updateProfileBiz) UpdateProfile(ctx context.Context, id int, data *usermodel.UserUpdate) error {

	if err := biz.store.UpdateProfile(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
