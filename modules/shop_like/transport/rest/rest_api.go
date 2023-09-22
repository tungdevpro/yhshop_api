package rest

import (
	"coffee_api/commons"
	shoplike "coffee_api/modules/shop_like"
	"coffee_api/modules/shop_like/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz shoplike.Business
}

func NewApi(biz shoplike.Business) shoplike.API {
	return &api{
		biz: biz,
	}
}

func (api *api) GetShopLikesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) GetLikedUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idShop := ctx.Param("id")
		id, err := strconv.Atoi(idShop)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		var params struct {
			commons.Paging
		}

		if err := ctx.ShouldBind(&params); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		items, err := api.biz.GetLikedUsers(ctx.Request.Context(), &entity.Filter{
			ShopId: id,
		}, &params.Paging)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, commons.NewSuccessResp(items, params.Paging, nil))
	}
}

func (api *api) CreateLikesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
