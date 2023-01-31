package restaurantservice

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/middleware"
	"golang_01/modules/restaurant/transport/gin"
	restaurantlikegin "golang_01/modules/restaurantlike/transport"
)

func RestaurantService(appCtx component.AppContext, router *gin.RouterGroup) error {
	restaurants := router.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))
		restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
		restaurants.GET("/:restaurant_id", restaurantgin.FindRestaurant(appCtx))
		restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurant(appCtx))
		restaurants.GET("/:restaurant_id/liked-users", restaurantlikegin.ListUserLikeRestaurant(appCtx))
	}
	return nil
}
