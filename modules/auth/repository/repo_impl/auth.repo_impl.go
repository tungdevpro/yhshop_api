package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/middleware"
	"coffee_api/modules/auth"
	authEntity "coffee_api/modules/auth/entity"
	userEntity "coffee_api/modules/user/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type authRepoImpl struct {
	appCtx commons.AppContext
}

func NewAuthRepoImpl(appCtx commons.AppContext) auth.Repository {
	return &authRepoImpl{
		appCtx: appCtx,
	}
}

func (r *authRepoImpl) Register(ctx context.Context, req *authEntity.RegisterDTO) error {
	r.appCtx.L.Lock()
	defer r.appCtx.L.Unlock()

	user := userEntity.User{
		Email: req.Email,
	}
	db := r.appCtx.GetDB()
	result := db.Where(&user).First(&user)

	if result.Error != nil || result.RowsAffected == 0 {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			accessToken, err := middleware.GenToken(r.appCtx.Cfg, middleware.JwtPayload{
				Email: user.Email,
				Role:  string(userEntity.Member),
			})
			if err != nil {
				return err
			}
			user.AccessToken = accessToken
			if err := db.Create(&user).Error; err != nil {
				return err
			}
			return nil
		}
		return result.Error
	}

	if user.Id != 0 {
		return errors.New(commons.ErrUserIsExist)
	}

	return nil
}

func (r *authRepoImpl) Login(ctx context.Context, loginDto *authEntity.LoginDTO) error {
	r.appCtx.L.Lock()
	defer r.appCtx.L.Unlock()

	user := userEntity.User{
		Email: loginDto.Email,
	}

	db := r.appCtx.GetDB()
	result := db.Where(&user).First(&user)

	if result.Error != nil || result.RowsAffected == 0 {
		return authEntity.ErrUnauthorized
	}

	if err := user.VerifyPassword(loginDto.Password); err != nil {
		return authEntity.ErrUnauthorized
	}

	return nil
}
