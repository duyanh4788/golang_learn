package admingin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/admin/biz"
	"golang_01/modules/user/model"
	"golang_01/modules/user/store"
	"net/http"
)

func UpdateStatusUserByAdmin(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userStatus usermodel.UserUpdateStatus
		if err := c.ShouldBind(&userStatus); err == nil {
			if userStatus.Status > 1 || userStatus.Status < 0 {
				panic(common.ErrInvalidRequest(errors.New("request status invalid")))
			}
		} else {
			panic(err)
		}

		user := c.MustGet(common.CurrentUser).(*usermodel.User)

		if userStatus.UserId == int(user.GetUserId()) {
			panic(usermodel.ErrUpdateStatus)
		}

		store := userstore.NewSqlStore(appContext.GetMainDBConnect())
		biz := adminbiz.NewUpdateStatusUserBiz(store)

		if err := biz.UpdateStatusUserByAdmin(c.Request.Context(), userStatus); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "update status success"))
	}
}
