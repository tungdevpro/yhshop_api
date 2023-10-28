package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/user"
	"coffee_api/modules/user/entity"
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
func (impl *userRepoImpl) ChangeVerifyEmail(ctx context.Context, email string) error {
	db := impl.appCtx.GetDB()

	result := db.Table(entity.User{}.TableName()).Where("email = ?", email).Update("is_email_verified", 1)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
