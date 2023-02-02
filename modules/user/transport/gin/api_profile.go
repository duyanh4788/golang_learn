package usergin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	usermodel "golang_01/modules/user/model"
	"net/http"
)

func Profile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		profile := c.MustGet(common.CurrentUser).(*usermodel.User)

		profile.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(profile, "success"))
	}
}
