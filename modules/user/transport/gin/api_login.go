package usergin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/component/hasher"
	"golang_01/component/tokenprovider/jwt"
	"golang_01/modules/user/biz"
	"golang_01/modules/user/model"
	"golang_01/modules/user/store"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnect()
		tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())

		store := userstore.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*20*30)

		account, err := biz.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account, "login success"))
	}
}
