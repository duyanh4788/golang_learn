package main

import (
	"demo_golang_02/component"
	"demo_golang_02/modules/restaurant/restauranttransport/ginrestaurant"
	"github.com/gin-gonic/gin"
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

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
			restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
		}
	}
	router.Run(":3010")
}
