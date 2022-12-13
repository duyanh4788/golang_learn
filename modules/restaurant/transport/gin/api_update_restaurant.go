package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/common"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		id, err := strconv.Atoi(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrIntenval(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err))
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "update success"))
	}
}
