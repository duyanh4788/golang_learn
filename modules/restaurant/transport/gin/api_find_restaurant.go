package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/biz"
	restaurantrepo "golang_01/modules/restaurant/repo"
	"golang_01/modules/restaurant/storage"
	restaurantlikestorage "golang_01/modules/restaurantlike/storage"
	"net/http"
)

func FindRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrIntenval(err))
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		likeStore := restaurantlikestorage.NewSqlStore(appContext.GetMainDBConnect())
		restaurantRepo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewFindRestaurantBiz(store, restaurantRepo)
		data, err := biz.FindRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data, ""))
	}
}
