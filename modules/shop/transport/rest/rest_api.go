package rest

import (
	"coffee_api/modules/shop"

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
	return func(ctx *gin.Context) {}
}

func (api *api) GetShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) CreateShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) DeleteShopHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
