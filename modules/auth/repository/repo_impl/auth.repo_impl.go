package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/middleware"
	"coffee_api/modules/auth"
	authEntity "coffee_api/modules/auth/entity"
	userEntity "coffee_api/modules/user/entity"
	"context"
	"errors"
)

type authRepoImpl struct {
	appCtx commons.AppContext
}

func NewAuthRepoImpl(appCtx commons.AppContext) auth.Repository {
	return &authRepoImpl{
		appCtx: appCtx,
	}
}

func (r *authRepoImpl) Register(ctx context.Context, req *authEntity.RegisterRequest) error {
	r.appCtx.L.Lock()
	defer r.appCtx.L.Unlock()

	user := userEntity.User{
		Email:    req.Email,
		FullName: req.FullName,
		Status:   userEntity.ACTIVE,
	}
	db := r.appCtx.GetDB()

	if err := db.First(&user, "email = ?", user.Email).Error; err != nil {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
		return err
	}

	if user.Id != 0 {
		return errors.New(commons.ErrUserIsExist)
	}

	token, err := middleware.GenToken(r.appCtx.Cfg, middleware.JwtPayload{
		UserId: user.Id,
		Role:   "Member",
	})
	if err != nil {
		return err
	}

	req.Token = token

	return nil
}

func (repo *authRepoImpl) Login(ctx context.Context, req *authEntity.LoginRequest) error {
	return nil
}
