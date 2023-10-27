package repoimpl

import (
	authEntity "coffee_api/modules/auth/entity"
	userEntity "coffee_api/modules/user/entity"
	"context"
)

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
		return nil, authEntity.ErrVerifiedYourAccount
	}

	resp := authEntity.LoginResponse{
		Id:              user.Id,
		FullName:        user.FullName,
		AccessToken:     user.AccessToken,
		IsEmailVerified: user.IsEmailVerified,
	}

	return &resp, nil
}
