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

func (r *authRepoImpl) Register(ctx context.Context, req *authEntity.RegisterDTO) (string, error) {
	r.appCtx.L.Lock()
	defer r.appCtx.L.Unlock()

	db := r.appCtx.GetDB()
	db.Begin()

	user := userEntity.User{
		Email: req.Email,
	}
	result := db.Where(&user).First(&user)

	if result.Error != nil || result.RowsAffected == 0 {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			accessToken, err := middleware.GenToken(r.appCtx.Cfg, middleware.JwtPayload{
				Email: user.Email,
				Role:  string(userEntity.Member),
			})
			if err != nil {
				db.Rollback()
				return "", err
			}
			user.AccessToken = accessToken
			user.Email = req.Email
			user.FullName = req.FullName
			user.Password = req.Password
			if err := db.Create(&user).Error; err != nil {
				db.Rollback()
				return "", err
			}
			return user.Uid.EncodeString(), nil
		}
		return "", result.Error
	}

	if user.Id != 0 {
		db.Rollback()
		return "", errors.New(commons.ErrUserIsExist)
	}

	db.Commit()
	return user.Uid.EncodeString(), nil
}

func (r *authRepoImpl) Login(ctx context.Context, loginDto *authEntity.LoginDTO) (*authEntity.LoginResponse, error) {
	r.appCtx.L.Lock()
	defer r.appCtx.L.Unlock()

	user := userEntity.User{
		Email: loginDto.Email,
	}

	db := r.appCtx.GetDB()
	result := db.Where(&user).First(&user)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, authEntity.ErrUnauthorized
	}

	if err := user.VerifyPassword(loginDto.Password); err != nil {
		return nil, authEntity.ErrUnauthorized
	}

	resp := authEntity.LoginResponse{
		Id:          user.Id,
		FullName:    user.FullName,
		AccessToken: user.AccessToken,
	}

	return &resp, nil
}
