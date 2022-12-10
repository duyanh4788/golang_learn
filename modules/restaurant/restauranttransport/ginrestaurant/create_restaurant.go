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

func CreateRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSqlStore(appContext.GetMainDBConnect())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data, ""))
	}
}

//type fakeCreateStore struct{}
//
//func (fakeCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
//	data.Id = 100
//	return nil
//}
