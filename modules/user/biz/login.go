package userbiz

import (
	"context"
	"golang_01/common"
	"golang_01/component/tokenprovider"
	"golang_01/modules/user/model"
)

type LoginStorage interface {
	FindUserByEmail(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type LoginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hashes        Hasher
	expiry        int
}

func NewLoginBiz(storeUser LoginStorage, toKenProvider tokenprovider.Provider, hashes Hasher, expiry int) *LoginBiz {
	return &LoginBiz{
		storeUser:     storeUser,
		tokenProvider: toKenProvider,
		hashes:        hashes,
		expiry:        expiry,
	}
}

func (biz *LoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUserByEmail(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPassWordInvalid
	}

	passHashed := biz.hashes.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPassWordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role.String(),
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrIntenval(err)
	}

	return accessToken, nil
}
