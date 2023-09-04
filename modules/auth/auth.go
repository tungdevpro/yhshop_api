package auth

import (
	"coffee_api/modules/auth/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	Register(context.Context, *entity.RegisterRequest) error
	Login(context.Context, *entity.LoginRequest) error
}

type Repository interface {
	Login(context.Context, *entity.LoginRequest) error
	Register(context.Context, *entity.RegisterRequest) error
}

type API interface {
	RegisterHandler() gin.HandlerFunc
	LoginHandler() gin.HandlerFunc
}
