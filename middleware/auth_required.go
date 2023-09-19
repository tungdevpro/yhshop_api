package middleware

import (
	"coffee_api/commons"
	"coffee_api/helpers"
	jwtexplore "coffee_api/middleware/jwt_explore"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indrasaputra/hashids"
)

func AuthRequired(appCtx commons.AppContext, bizJwt jwtexplore.Business) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := helpers.GetTokenFromAuthHeader(ctx.GetHeader("Authorization"))

		uid, err := Validate(appCtx.Cfg, tokenString)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		id, err := hashids.DecodeHash([]byte(uid))
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		user, err := bizJwt.FindUser(ctx.Request.Context(), int(id))
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		fmt.Println("log: ", user)

		ctx.Next()
	}
}
