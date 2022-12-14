package main

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/uploadprovider"
	"golang_01/middleware"
	"golang_01/modules/restaurant/transport/gin"
	"golang_01/modules/upload/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dns := os.Getenv("MYSQL")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("s3Region")
	s3Apikey := os.Getenv("s3Apikey")
	s3SecretKey := os.Getenv("s3SecretKey")
	s3SDomain := os.Getenv("s3SDomain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3Apikey, s3SecretKey, s3SDomain)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Can not connect DB")
	}

	log.Println("Connect to", db)

	if err := runService(db, s3Provider); err != nil {
		log.Println(err)
	}

}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvide) error {
	appCtx := component.NewAppContext(db, upProvider)

	router := gin.Default()
	router.SetTrustedProxies([]string{"12.4.27.15"})
	router.Use(middleware.Recover(appCtx))

	router.POST("/upload", uploadgin.Upload(appCtx))

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))
			restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
			restaurants.GET("/:restaurant_id", restaurantgin.FindRestaurant(appCtx))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurant(appCtx))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurant(appCtx))
		}
	}
	return router.Run(":3010")
}
