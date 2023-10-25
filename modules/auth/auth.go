package auth

import (
	"context"

	"coffee_api/modules/auth/entity"

	"github.com/gin-gonic/gin"
)

type Business interface {
	Login(context.Context, *entity.LoginDTO) (*entity.LoginResponse, error)
	Register(context.Context, *entity.RegisterDTO) (*entity.RegisterReponse, error)
}

type Repository interface {
	Login(context.Context, *entity.LoginDTO) (*entity.LoginResponse, error)
	Register(context.Context, *entity.RegisterDTO) (*entity.RegisterReponse, error)
}

type API interface {
	LoginHandler() gin.HandlerFunc
	RegisterHandler() gin.HandlerFunc
	VerifyMail() gin.HandlerFunc
}
