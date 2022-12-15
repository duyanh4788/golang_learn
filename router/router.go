package mainrouter

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/uploadprovider"
	"golang_01/middleware"
	"golang_01/router/restaurant"
	"golang_01/router/uploadaws"
	"gorm.io/gorm"
)

func MainServices(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	router := gin.Default()
	router.SetTrustedProxies([]string{"12.4.27.15"})
	router.Use(middleware.Recover(appCtx))

	v1 := router.Group("/v1")

	if err := uploadawsservice.UploadService(appCtx, v1); err != nil {
		panic(err)
	}

	if err := restaurantservice.RestaurantService(appCtx, v1); err != nil {
		panic(err)
	}
	return router.Run(":3010")
}
