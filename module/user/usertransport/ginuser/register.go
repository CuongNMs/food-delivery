package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/component/hasher"
	"food-delivery/module/user/userbiz"
	"food-delivery/module/user/usermodel"
	"food-delivery/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		storage := userstorage.NewSQLStorage(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(storage, md5)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
