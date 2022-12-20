package userbiz

import (
	"context"
	"golang_01/common"
	"golang_01/modules/user/model"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hashes interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStore RegisterStore
	hashes        Hashes
}

func NewRegisterBiz(store RegisterStore, hashes Hashes) *registerBiz {
	return &registerBiz{
		registerStore: store,
		hashes:        hashes,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExited
	}

	salt := common.GenSalt(50)

	data.Password = biz.hashes.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
