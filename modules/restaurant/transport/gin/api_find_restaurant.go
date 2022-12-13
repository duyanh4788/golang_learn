package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/common"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func FindRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("restaurant_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		data, err := biz.FindRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data, ""))
	}
}
