package rest

import (
	verifymail "coffee_api/modules/verify_mail"
	"context"

	"github.com/gin-gonic/gin"
)

type api struct {
}

func NewApi() verifymail.API {
	return &api{}
}

func (api *api) CreateMailHandler(context context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
