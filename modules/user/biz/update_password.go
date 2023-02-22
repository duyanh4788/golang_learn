package userbiz

import (
	"context"
	"golang_01/common"
	"golang_01/component/tokenprovider"
	"golang_01/modules/user/model"
)

type UpdatePassWordStore interface {
	UpdatePassWord(ctx context.Context, id int, data *usermodel.UserUpdatePassWord) error
}

type updatePassWordBiz struct {
	store         UpdatePassWordStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
}

func NewUpdatePassWordBiz(store UpdatePassWordStore, tokenProvider tokenprovider.Provider, hasher Hasher) *updatePassWordBiz {
	return &updatePassWordBiz{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
	}
}

func (biz *updatePassWordBiz) UpdatePassWord(ctx context.Context, user *usermodel.User, data *usermodel.UserUpdatePassWord) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	passOldHashed := biz.hasher.Hash(data.OldPassword + user.Salt)

	if user.Password != passOldHashed {
		return usermodel.ErrPassWrong
	}
	salt := common.GenSalt(50)
	data.NewPassword = biz.hasher.Hash(data.NewPassword + salt)
	data.Salt = salt

	if err := biz.store.UpdatePassWord(ctx, user.Id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
