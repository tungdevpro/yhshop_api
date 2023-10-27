package auth

import (
	"context"

	"coffee_api/modules/auth/entity"

	"github.com/gin-gonic/gin"
)

type Business interface {
	Login(context.Context, *entity.LoginDTO) (*entity.LoginResponse, error)
	Register(context.Context, *entity.RegisterDTO) (*entity.RegisterReponse, error)
	VerifyOTP(context.Context, *entity.OTPRequest) (bool, error)
}

type Repository interface {
	Login(context.Context, *entity.LoginDTO) (*entity.LoginResponse, error)
	Register(context.Context, *entity.RegisterDTO) (*entity.RegisterReponse, error)
	VerifyOTP(context.Context, *entity.OTPRequest) (bool, error)
}

type API interface {
	LoginHandler() gin.HandlerFunc
	RegisterHandler() gin.HandlerFunc
	VerifyOTPHandler() gin.HandlerFunc
}
