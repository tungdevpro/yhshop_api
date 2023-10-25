package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/helpers"
	"coffee_api/middleware"
	"coffee_api/modules/auth"
	"coffee_api/modules/auth/entity"
	authEntity "coffee_api/modules/auth/entity"
	userEntity "coffee_api/modules/user/entity"
	"context"
	"errors"
	"fmt"

	"github.com/indrasaputra/hashids"
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

func (r *authRepoImpl) Register(ctx context.Context, req *authEntity.RegisterDTO) (*entity.RegisterReponse, error) {
	r.appCtx.L.Lock()
	defer r.appCtx.L.Unlock()

	db := r.appCtx.GetDB()
	db.Begin()

	doc := entity.CreateUser{
		Email: req.Email,
	}
	result := db.Where(&doc).First(&doc)

	if result.Error != nil || result.RowsAffected == 0 {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			doc.Email = req.Email
			doc.Password = req.Password
			doc.FullName = req.FullName
			if err := db.Create(&doc).Error; err != nil {
				db.Rollback()
				return nil, err
			}
			pId := hashids.ID(doc.Id)
			uid, _ := hashids.EncodeID(pId)
			accessToken, err := middleware.GenToken(r.appCtx.Cfg, middleware.JwtPayload{
				Id:   string(uid),
				Role: string(commons.Member),
			})
			if err != nil {
				db.Rollback()
				return nil, err
			}

			resp := entity.RegisterReponse{
				Id:          doc.Id,
				Uid:         string(uid),
				AccessToken: accessToken,
				Email:       doc.Email,
				FullName:    doc.FullName,
			}

			otp := helpers.EncodeToString(5)
			subject := "Account authentication"
			content := fmt.Sprintf(`
			<h1>Hi! %s</h1>
			<p>
			Thank you for choosing %s. Use the following OTP to complete your Sign Up procedures. OTP is valid for 5 minutes
			</p>
			<h2 style="background: #00466a;margin: 0 auto;width: max-content;padding: 0 10px;color: #fff;border-radius: 4px;">%s</h2>
			`, doc.FullName, r.appCtx.Cfg.ApplicationName, otp)
			to := []string{doc.Email}

			err = r.appCtx.Mailer.SendEmail(subject, content, to, nil, nil, nil)
			if err != nil {
				return nil, err
			}

			db.Commit()
			return &resp, nil
		}
		return nil, result.Error
	}

	if doc.Id != 0 {
		db.Rollback()
		return nil, errors.New(commons.ErrUserIsExist)
	}

	db.Commit()
	return nil, nil
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
