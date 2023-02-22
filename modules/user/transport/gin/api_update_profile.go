package usergin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/user/biz"
	"golang_01/modules/user/model"
	"golang_01/modules/user/store"
	"net/http"
)

func UpdateProfile(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotUpdateEntity(usermodel.EntityName, err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if int(requester.GetStatus()) == 0 {
			panic(common.ErrDisableStatus(usermodel.EntityName, string(requester.GetEmail()), nil))
		}

		store := userstore.NewSqlStore(appContext.GetMainDBConnect())
		biz := userbiz.NewUpdateProfileBiz(store)

		if err := biz.UpdateProfile(c.Request.Context(), requester.GetUserId(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "update success"))
	}
}
