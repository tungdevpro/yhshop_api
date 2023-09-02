package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/auth/entity"
	"coffee_api/modules/auth/repository"
	"context"
)

type authRepoImpl struct {
	appCtx commons.AppContext
}

func NewAuthRepoImpl(appCtx commons.AppContext) repository.Repository {
	return &authRepoImpl{
		appCtx: appCtx,
	}
}

func (repo *authRepoImpl) Register(ctx context.Context, request entity.RegisterRequest) {

}
func (repo *authRepoImpl) Login(ctx context.Context, request entity.LoginRequest) {}
