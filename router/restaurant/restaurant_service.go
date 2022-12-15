package restaurantservice

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/modules/restaurant/transport/gin"
)

func RestaurantService(appCtx component.AppContext, router *gin.RouterGroup) error {
	restaurants := router.Group("/restaurants")
	{
		restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))
		restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
		restaurants.GET("/:restaurant_id", restaurantgin.FindRestaurant(appCtx))
		restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurant(appCtx))
	}
	return nil
}
