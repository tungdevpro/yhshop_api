package middleware

import (
	"coffee_api/commons"
	"coffee_api/helpers"

	"github.com/gin-gonic/gin"
)

func AuthRequired(appCtx commons.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := helpers.ExtraTokenFromHeader(ctx.GetHeader("Authorization"))

		Validate(appCtx.Cfg, tokenString)
		ctx.Next()
	}
}
