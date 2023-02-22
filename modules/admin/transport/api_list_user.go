package admingin

import (
	"github.com/gin-gonic/gin"
	"golang_01/common"
	"golang_01/component"
	adminbiz "golang_01/modules/admin/biz"
	adminmodel "golang_01/modules/admin/model"
	userstore "golang_01/modules/user/store"
	"net/http"
)

func GetListUsersByAdmin(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter adminmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.Fulfill()

		store := userstore.NewSqlStore(appContext.GetMainDBConnect())
		biz := adminbiz.NewListUsersByAdminBiz(store)

		result, err := biz.ListUsersByAdmin(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(true)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter, "success", ""))
	}
}
