package restaurants

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	GetListRestaurant(context.Context)
	GetRestaurantById(context.Context)
}

type Repository interface {
	GetListRestaurant(context.Context)
	GetRestaurantById(context.Context)
}

type API interface {
	ListRestaurantHandler() gin.HandlerFunc
	GetRestaurantHandler() gin.HandlerFunc
}
