package rest

import (
	"coffee_api/commons"
	"coffee_api/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz user.Business
}

func NewApi(biz user.Business) user.API {
	return &api{biz: biz}
}

func (api *api) GetProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := ctx.MustGet(commons.CurrentUser).(commons.Requester)

		ctx.JSON(http.StatusOK, commons.SimpleSuccessResp(data.GetEmail()))
	}
}

func (api *api) DeleteUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *api) UpdateProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
