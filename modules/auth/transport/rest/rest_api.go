package rest

import (
	"coffee_api/modules/auth/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz business.IBusiness
}

func NewApi(biz business.IBusiness) *api {
	return &api{
		biz: biz,
	}
}

func (api *api) RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "register",
		})
	}
}
func (api *api) LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "login",
		})
	}
}
