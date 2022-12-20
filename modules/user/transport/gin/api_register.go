package usergin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/component/hasher"
	"golang_01/modules/user/biz"
	"golang_01/modules/user/model"
	"golang_01/modules/user/store"
	"net/http"
)

func Register(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnect()
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "register success"))
	}
}
