package user

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetProfile(context.Context)
	DelUser(context.Context)
	UpdateProfile(context.Context)
	ChangeVerifyEmail(context.Context, string) error
}

type Business interface {
	GetProfile(context.Context)
	DelUser(context.Context)
	UpdateProfile(context.Context)
}

type API interface {
	GetProfileHandler() gin.HandlerFunc
	DeleteUserHandler() gin.HandlerFunc
	UpdateProfileHandler() gin.HandlerFunc
}
