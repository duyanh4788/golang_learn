package usergin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"net/http"
)

func Profile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		profile := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(profile, "success"))
	}
}
