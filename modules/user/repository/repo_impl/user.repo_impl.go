package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/user"
	"context"
)

type userRepoImpl struct {
	appCtx commons.AppContext
}

func NewUserRepoImpl(appCtx commons.AppContext) user.Repository {
	return &userRepoImpl{
		appCtx: appCtx,
	}
}

func (impl *userRepoImpl) GetProfile(ctx context.Context)    {}
func (impl *userRepoImpl) DelUser(ctx context.Context)       {}
func (impl *userRepoImpl) UpdateProfile(ctx context.Context) {}
