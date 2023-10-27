package rest

import (
	"net/http"

	"coffee_api/commons"
	"coffee_api/modules/auth"
	"coffee_api/modules/auth/entity"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz auth.Business
}

func NewApi(biz auth.Business) auth.API {
	return &api{
		biz: biz,
	}
}

func (api *api) RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data entity.RegisterDTO

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		result, err := api.biz.Register(ctx.Request.Context(), &data)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, commons.CreateNewSuccessResp(result.Uid, entity.MsgCreateNewUser))
	}
}

func (api *api) LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto entity.LoginDTO

		if err := ctx.ShouldBind(&dto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		result, err := api.biz.Login(ctx.Request.Context(), &dto)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}

		if !result.IsEmailVerified {
			ctx.AbortWithStatusJSON(http.StatusForbidden, commons.NewAppError(int(entity.NotVerified), entity.ErrVerifiedYourAccount.Error()))
			return
		}

		ctx.JSON(http.StatusOK, commons.CreateNewSuccessResp(result, entity.MsgLoginSuccess))
	}
}

func (api *api) VerifyOTPHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto = entity.OTPRequest{}
		if err := ctx.ShouldBind(&dto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}
		result, err := api.biz.VerifyOTP(ctx, &dto)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, commons.NewAppError(-1, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, commons.CreateNewSuccessResp(result, entity.MsgLoginSuccess))

	}
}
