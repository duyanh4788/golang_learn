package restaurantgin

import (
	"github.com/gin-gonic/gin"
	common2 "golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("restaurant_id"))

		if err != nil {
			panic(common2.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common2.SimpleSuccessResponse(true, "delete success"))
	}
}
