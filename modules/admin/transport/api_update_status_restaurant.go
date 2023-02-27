package admingin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/admin/biz"
	"golang_01/modules/restaurant/model"
	"golang_01/modules/restaurant/storage"
	"net/http"
)

func UpdateStatusRestaurantByAdmin(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var restaurantStatus restaurantmodel.UpdateStatusRestaurant

		if err := c.ShouldBind(&restaurantStatus); err == nil {
			if restaurantStatus.Status > 1 || restaurantStatus.Status < 0 {
				panic(common.ErrInvalidRequest(errors.New("request status invalid")))
			}
		} else {
			panic(err)
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := adminbiz.NewUpdateStatusRestaurantBiz(store)

		if err := biz.UpdateRestaurantByAdmin(c.Request.Context(), &restaurantStatus); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true, "update status success"))
	}
}
