package auth

import (
	"coffee_api/modules/auth/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	Register(context.Context, *entity.RegisterDTO) (string, error)
	Login(context.Context, *entity.LoginDTO) (*entity.LoginResponse, error)
}

type Repository interface {
	Login(context.Context, *entity.LoginDTO) (*entity.LoginResponse, error)
	Register(context.Context, *entity.RegisterDTO) (string, error)
}

type API interface {
	RegisterHandler() gin.HandlerFunc
	LoginHandler() gin.HandlerFunc
}
