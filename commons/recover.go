package commons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover(appCtx AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")
				ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
				panic(err)
			}
		}()

		ctx.Next()
	}
}
