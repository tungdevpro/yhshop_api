package verifymail

import (
	"coffee_api/modules/verify_mail/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type API interface {
	CreateMail(context.Context) gin.HandlerFunc
}

type Business interface {
	CreateMail(context.Context, entity.VerifyMail) error
}

type Repository interface {
	CreateMail(context.Context, entity.VerifyMail) error
}
