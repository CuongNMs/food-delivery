package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	bizrestaurant "food-delivery/module/restaurant/biz"
	"food-delivery/module/restaurant/model"
	storagerestaurant "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data []restaurantmodel.Restaurant

		storage := storagerestaurant.NewSQLStorage(db)
		biz := bizrestaurant.NewListRestaurantBiz(storage)
		data, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask()
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, pagingData, filter))
	}
}
