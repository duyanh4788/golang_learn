package userservice

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/middleware"
	"golang_01/modules/user/transport/gin"
	"net/http"
)

func UserService(appCtx component.AppContext, router *gin.RouterGroup) error {
	user := router.Group("/user")
	{
		user.POST("/register", usergin.Register(appCtx))
		user.POST("/login", usergin.Login(appCtx))
		user.GET("/profile", middleware.RequireAuth(appCtx), usergin.Profile(appCtx))
		user.GET("/admin", middleware.RequireAuth(appCtx), middleware.RequireRole(appCtx, "mod"),
			func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{"data": "ok"})
			})
	}
	return nil
}
