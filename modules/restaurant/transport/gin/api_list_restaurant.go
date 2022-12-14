package restaurantgin

import (
	"github.com/gin-gonic/gin"
	common2 "golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurant/storage"
	"net/http"
)

func ListRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common2.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.Fulfill()

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common2.NewSuccessResponse(result, paging, filter, "success", ""))
	}
}
