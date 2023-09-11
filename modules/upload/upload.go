package upload

import (
	"coffee_api/modules/upload/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	UploadFile(context.Context, *entity.UploadDTO) error
}

type Repository interface {
	UploadFile(context.Context, *entity.UploadDTO) error
}

type API interface {
	UploadFileHandler() gin.HandlerFunc
}
