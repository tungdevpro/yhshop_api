package rest

import (
	"coffee_api/commons"
	shoplike "coffee_api/modules/shop_like"
	"coffee_api/modules/shop_like/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indrasaputra/hashids"
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

func (api *api) CreateUserLikeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		if idParam == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, entity.EmptyParamIdShop.Error()))
			return
		}

		hashId, err := hashids.DecodeHash([]byte(string(idParam)))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		requester := ctx.MustGet(commons.CurrentUser).(commons.Requester)

		data := entity.ShopLike{
			ShopId: int(hashId),
			UserId: requester.GetUserId(),
		}

		id, err := api.biz.CreateUserLike(ctx.Request.Context(), data.UserId, data.ShopId)

		if err != nil {
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
				return
			}
		}

		ctx.JSON(http.StatusOK, commons.SimpleSuccessResp(id))
	}
}

func (api *api) DeleteUserLikeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		if idParam == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, entity.EmptyParamIdShop.Error()))
			return
		}

		hashId, err := hashids.DecodeHash([]byte(string(idParam)))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		requester := ctx.MustGet(commons.CurrentUser).(commons.Requester)

		data := entity.ShopLike{
			ShopId: int(hashId),
			UserId: requester.GetUserId(),
		}

		if err := api.biz.DeleteUserLike(ctx.Request.Context(), data.UserId, data.ShopId); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, commons.SimpleSuccessResp(true))
	}
}
