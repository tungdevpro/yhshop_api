package rest

import (
	shoplike "coffee_api/modules/shop_like"

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

func (api *api) GetShopLikes() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) GetLikedUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
