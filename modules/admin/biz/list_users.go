package adminbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/admin/model"
	"golang_01/modules/user/model"
)

type ListUsersByAdminStore interface {
	ListUsersByAdmin(
		ctx context.Context,
		filter *adminmodel.Filter,
		paging *common.Paging,
	) ([]usermodel.User, error)
}

type listUsersByAdminBiz struct {
	store ListUsersByAdminStore
}

func NewListUsersByAdminBiz(store ListUsersByAdminStore) *listUsersByAdminBiz {
	return &listUsersByAdminBiz{store: store}
}

func (biz *listUsersByAdminBiz) ListUsersByAdmin(
	ctx context.Context,
	filter *adminmodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {

	result, err := biz.store.ListUsersByAdmin(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
