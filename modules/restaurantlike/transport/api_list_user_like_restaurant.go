package restaurantlikegin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurantlike/biz"
	"golang_01/modules/restaurantlike/model"
	"golang_01/modules/restaurantlike/storage"
	"net/http"
)

func ListUserLikeRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.Fulfill()

		store := restaurantlikestore.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUserLikeRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter, "success", ""))

	}
}
