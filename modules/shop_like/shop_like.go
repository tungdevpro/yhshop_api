package shoplike

import (
	"context"

	"github.com/gin-gonic/gin"
)

type API interface {
	GetShopLikes() gin.HandlerFunc
}

type Business interface {
	GetShopLikes(context.Context, []int) (map[int]int, error)
}

type Repository interface {
	GetShopLikes(context.Context, []int) (map[int]int, error)
}
