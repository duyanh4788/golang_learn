package main

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/middleware"
	"golang_01/modules/restaurant/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dns := os.Getenv("MYSQL")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Can not connect DB")
	}

	log.Println("Connect to", db)

	appContext := component.NewAppContext(db)

	router := gin.Default()
	router.SetTrustedProxies([]string{"12.4.27.15"})
	router.Use(middleware.Recover(appContext))

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurant(appContext))
			restaurants.GET("", restaurantgin.ListRestaurant(appContext))
			restaurants.GET("/:restaurant_id", restaurantgin.FindRestaurant(appContext))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurant(appContext))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurant(appContext))
		}
	}
	router.Run(":3010")
}
