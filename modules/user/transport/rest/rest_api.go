package rest

import (
	"coffee_api/modules/user"

	"github.com/gin-gonic/gin"
)

type api struct {
}

func NewApi() user.API {
	return &api{}
}

func (api *api) GetProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) DeleteUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) UpdateProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
