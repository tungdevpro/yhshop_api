package middleware

import (
	"coffee_api/commons"
	"coffee_api/configs/prefix"
	"coffee_api/helpers"
	jwtexplore "coffee_api/middleware/jwt_explore"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/indrasaputra/hashids"
)

var (
	AllowPaths = []string{
		fmt.Sprintf("%s%s%s", prefix.V1, prefix.Auth, prefix.Register),
		fmt.Sprintf("%s%s%s", prefix.V1, prefix.Auth, prefix.Login),
		fmt.Sprintf("%s%s%s", prefix.V1, prefix.Auth, prefix.VerifyOTP),
	}
)

func AuthRequired(appCtx commons.AppContext, bizJwt jwtexplore.Business) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if slices.Contains(AllowPaths, ctx.FullPath()) {
			ctx.Next()
			return
		}
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

		if ok := user.IsActive(); !ok {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New(commons.ErrUserIsInActive))
			return
		}

		ctx.Set(commons.CurrentUser, user)
		ctx.Next()
	}
}
