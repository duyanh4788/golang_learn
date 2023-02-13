package restaurantlikegin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/storage"
	"golang_01/modules/restaurantlike/biz"
	"golang_01/modules/restaurantlike/model"
	"golang_01/modules/restaurantlike/storage"
	"net/http"
)

func UserUnLikeRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		restaurantStore := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		restaurantBiz := restaurantbiz.NewFindRestaurantBiz(restaurantStore, nil)

		restaurant, err := restaurantBiz.FindRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.RestaurantLike{
			RestaurantId: restaurant.Id,
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantlikebiz.NewUserUnLikedRestaurantBiz(store, restaurantStore)

		mesage, err := biz.UserUnLikedRestaurant(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(true, nil, nil, "success", mesage))

	}
}
