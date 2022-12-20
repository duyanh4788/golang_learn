package middleware

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
)

func RequireRole(appCtc component.AppContext, roles ...string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for i := range roles {
			if u.GetRole() == roles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrPermission(u.GetEmail(), nil))
	}
}
