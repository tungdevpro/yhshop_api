package shoplike

import (
	"coffee_api/commons"
	"coffee_api/modules/shop_like/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type API interface {
	GetShopLikes() gin.HandlerFunc
	GetLikedUsers() gin.HandlerFunc
}

type Business interface {
	GetShopLikes(context.Context, []int) (map[int]int, error)
	GetLikedUsers(context.Context, *entity.Filter) ([]commons.SimpleUser, error)
}

type Repository interface {
	GetShopLikes(context.Context, []int) (map[int]int, error)
	GetLikedUsers(context.Context, *entity.Filter) ([]commons.SimpleUser, error)
}
