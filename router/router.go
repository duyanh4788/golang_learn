package mainrouter

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/uploadprovider"
	"golang_01/middleware"
	"golang_01/pubsub/pubsublocal"
	adminservice "golang_01/router/admin"
	"golang_01/router/restaurant"
	"golang_01/router/uploadaws"
	userservice "golang_01/router/user"
	"golang_01/subscriber"
	"gorm.io/gorm"
	"log"
)

func MainServices(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContext(db, upProvider, secretKey, pubsublocal.NewPubSub())
	//subscriber.Setup(appCtx)
	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatal(err)
	}
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

	if err := userservice.UserService(appCtx, v1); err != nil {
		panic(err)
	}

	if err := adminservice.AdminService(appCtx, v1); err != nil {
		panic(err)
	}

	return router.Run(":3010")
}
