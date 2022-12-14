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

func CreateRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common2.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common2.SimpleSuccessResponse(data, ""))
	}
}
