package category

import (
	"context"

	"github.com/gin-gonic/gin"
)

type API interface {
	GetListCategories() gin.HandlerFunc
}

type Business interface {
	GetListCategories(context.Context)
}

type Repository interface {
	GetListCategories(context.Context)
}
