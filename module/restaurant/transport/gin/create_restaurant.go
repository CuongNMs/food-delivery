package ginrestaurant

import (
	"fmt"
	"food-delivery/common"
	"food-delivery/component/appctx"
	bizrestaurant "food-delivery/module/restaurant/biz"
	"food-delivery/module/restaurant/model"
	storagerestaurant "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		db := appCtx.GetMainDBConnection()

		go func() {
			defer common.AppRecover()
			arr := []int{}
			fmt.Println(arr[0])
		}()

		if err := c.ShouldBind(&data); err != nil {
			c.JSONP(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := storagerestaurant.NewSQLStorage(db)
		biz := bizrestaurant.NewCreateRestaurantBiz(storage)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
