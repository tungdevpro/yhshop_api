package restaurants

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	GetListShop(context.Context)
	GetShopById(context.Context)
}

type Repository interface {
	GetListShop(context.Context)
	GetShopById(context.Context)
}

type API interface {
	ListShopHandler() gin.HandlerFunc
	GetShopHandler() gin.HandlerFunc
}
