package adminservice

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/middleware"
	"golang_01/modules/admin/transport"
)

func AdminService(appCtx component.AppContext, router *gin.RouterGroup) error {
	admin := router.Group("/admin", middleware.RequireAuth(appCtx), middleware.RequireRole(appCtx, "admin"))
	{
		admin.GET("/list-user", admingin.GetListUsersByAdmin(appCtx))
	}

	return nil
}
