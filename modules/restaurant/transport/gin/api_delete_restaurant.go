package restaurantgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/common"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("restaurant_id"))
		fmt.Println(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "delete success"))
	}
}
