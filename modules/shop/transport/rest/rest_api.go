package rest

import (
	"coffee_api/commons"
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz shop.Business
}

func NewApi(biz shop.Business) shop.API {
	return &api{
		biz: biz,
	}
}

func (api *api) ListShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter entity.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		// items, err := api.biz.GetListShop(ctx.Request.Context(), &filter)
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
		// 	return
		// }

		// fmt.Println("items: ", items)

		ctx.JSON(http.StatusOK, commons.SimpleSuccessResp(nil))
	}
}

func (api *api) GetShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) CreateShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto entity.CreateShopDTO

		if err := ctx.ShouldBind(&dto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		result, err := api.biz.CreateShop(ctx.Request.Context(), &dto)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, commons.SimpleSuccessResp(result))

	}
}

func (api *api) UpdateShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) DeleteShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
