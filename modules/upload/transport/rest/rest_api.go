package rest

import (
	"coffee_api/modules/upload"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz upload.Business
}

func NewApi(biz upload.Business) *api {
	return &api{biz: biz}
}

func (api *api) UploadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
