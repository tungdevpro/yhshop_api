package shop

import (
	"coffee_api/commons"
	"coffee_api/modules/shop/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type API interface {
	ListShopHandler() gin.HandlerFunc
	GetShopHandler() gin.HandlerFunc
	CreateShopHandler() gin.HandlerFunc
	UpdateShopHandler() gin.HandlerFunc
	DeleteShopHandler() gin.HandlerFunc
}

type Business interface {
	GetListShop(context.Context, *entity.Filter, *commons.Paging) ([]entity.Shop, error)
	GetShopById(context.Context, string) (*entity.Shop, error)
	CreateShop(context.Context, *entity.CreateShopDTO) (string, error)
	DeleteShop(context.Context, string) bool
	IncrementLikeCount(context.Context, int) error
	DecrementLikeCount(context.Context, int) error
}

type Repository interface {
	GetListShop(context.Context, *entity.Filter, *commons.Paging, ...string) ([]entity.Shop, error)
	GetShopById(context.Context, int) (*entity.Shop, error)
	CreateShop(context.Context, *entity.CreateShopDTO) (string, error)
	DeleteShop(context.Context, string) bool
	IncrementLikeCount(context.Context, int) error
	DecrementLikeCount(context.Context, int) error
}
