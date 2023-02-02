package restaurantlikegin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	"golang_01/modules/restaurant/biz"
	"golang_01/modules/restaurant/storage"
	"golang_01/modules/restaurantlike/biz"
	restaurantlikemodel "golang_01/modules/restaurantlike/model"
	"golang_01/modules/restaurantlike/storage"
	usermodel "golang_01/modules/user/model"
	"net/http"
)

func UserLikeRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantlikemodel.RestaurantLike

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

		userInfo := c.MustGet(common.CurrentUser).(*usermodel.User)

		data.RestaurantId = restaurant.Id
		data.UserId = userInfo.Id

		store := restaurantlikestorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store)

		mesage, err := biz.UserLikeRestaurant(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(true, nil, nil, "success", mesage))

	}
}
