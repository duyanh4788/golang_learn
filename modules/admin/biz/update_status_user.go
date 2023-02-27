package adminbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/user/model"
)

type UpdateStatusUserStore interface {
	UpdateStatusUser(ctx context.Context, userStatus *usermodel.UserUpdateStatus) error
}

type updateStatusUserBiz struct {
	store UpdateStatusUserStore
}

func NewUpdateStatusUserBiz(store UpdateStatusUserStore) *updateStatusUserBiz {
	return &updateStatusUserBiz{store}
}

func (biz *updateStatusUserBiz) UpdateStatusUserByAdmin(ctx context.Context, userStatus usermodel.UserUpdateStatus) error {
	if err := biz.store.UpdateStatusUser(ctx, &userStatus); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
