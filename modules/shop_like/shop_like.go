package shoplike

import (
	"coffee_api/commons"
	"coffee_api/modules/shop_like/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type API interface {
	GetShopLikesHandler() gin.HandlerFunc
	GetLikedUsersHandler() gin.HandlerFunc
	CreateLikesHandler() gin.HandlerFunc
}

type Business interface {
	GetShopLikes(context.Context, []int) (map[int]int, error)
	GetLikedUsers(context.Context, *entity.Filter, *commons.Paging) ([]commons.SimpleUser, error)
	CreateLikes(context.Context) (*string, error)
}

type Repository interface {
	GetShopLikes(context.Context, []int) (map[int]int, error)
	GetLikedUsers(context.Context, *entity.Filter, *commons.Paging) ([]commons.SimpleUser, error)
	CreateLikes(context.Context) (*string, error)
}
