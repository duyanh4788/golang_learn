package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/component/tokenprovider/jwt"
	"golang_01/modules/user/model"
	"strings"
)

type AuthenticaStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongHeader"),
	)
}

func ExtractTokenFromHeader(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext, store AuthenticaStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := ExtractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrDisableStatus(common.CurrentUser, user.Email, err))
		}
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
