package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/component/common"
	"golang_01/modules/restaurant/restaurantbiz"
	"golang_01/modules/restaurant/restaurantmodel"
	"golang_01/modules/restaurant/restaurantstorage"
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
