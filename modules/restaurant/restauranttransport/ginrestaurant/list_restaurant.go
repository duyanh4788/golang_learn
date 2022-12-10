package ginrestaurant

import (
	"demo_golang_02/component"
	"demo_golang_02/component/common"
	"demo_golang_02/modules/restaurant/restaurantbiz"
	"demo_golang_02/modules/restaurant/restaurantmodel"
	"demo_golang_02/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter, "success", ""))
	}
}
