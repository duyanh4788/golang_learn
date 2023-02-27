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

func UpdatePassWord(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserUpdatePassWord

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotUpdateEntity(usermodel.EntityName, err))
		}

		user := c.MustGet(common.CurrentUser).(*usermodel.User)

		tokenProvider := jwt.NewTokenJwtProvider(appContext.SecretKey())
		md5 := hasher.NewMd5Hash()

		store := userstore.NewSqlStore(appContext.GetMainDBConnect())
		biz := userbiz.NewUpdatePassWordBiz(store, tokenProvider, md5)

		if err := biz.UpdatePassWord(c.Request.Context(), user, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "update success"))
	}
}
