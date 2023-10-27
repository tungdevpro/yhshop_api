package repoimpl

import (
	"coffee_api/commons"
	verifymail "coffee_api/modules/verify_mail"
	"coffee_api/modules/verify_mail/entity"

	"context"
)

type verifyMail struct {
	appCtx commons.AppContext
}

func NewVerifyMail(appCtx commons.AppContext) verifymail.Repository {
	return &verifyMail{
		appCtx: appCtx,
	}
}

func (impl *verifyMail) CreateMail(context context.Context, param entity.VerifyMail) error {
	db := impl.appCtx.GetDB()

	doc := entity.VerifyMail{
		Email:      param.Email,
		FullName:   param.FullName,
		SecretCode: param.SecretCode,
	}

	if err := db.Create(&doc).Error; err != nil {
		return err
	}
	return nil
}

func (impl *verifyMail) CheckOTPMail(context context.Context, param entity.VerifyMail) error {
	db := impl.appCtx.GetDB()

	doc := entity.VerifyMail{
		Email:      param.Email,
		SecretCode: param.SecretCode,
	}

	result := db.Where(&doc).First(&doc)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
