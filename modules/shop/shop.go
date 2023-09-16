package shop

import (
	"coffee_api/commons"
	"coffee_api/modules/shop/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	GetListShop(context.Context, *entity.Filter, *commons.Paging) ([]entity.Shop, error)
	GetShopById(context.Context) (entity.Shop, error)
	CreateShop(context.Context, *entity.CreateShopDTO) (string, error)
	DeleteShop(context.Context)
}

type Repository interface {
	GetListShop(context.Context, *entity.Filter, *commons.Paging) ([]entity.Shop, error)
	GetShopById(context.Context) (entity.Shop, error)
	CreateShop(context.Context, *entity.CreateShopDTO) (string, error)
	DeleteShop(context.Context)
}

type API interface {
	ListShopHandler() gin.HandlerFunc
	GetShopHandler() gin.HandlerFunc
	CreateShopHandler() gin.HandlerFunc
	UpdateShopHandler() gin.HandlerFunc
	DeleteShopHandler() gin.HandlerFunc
}
