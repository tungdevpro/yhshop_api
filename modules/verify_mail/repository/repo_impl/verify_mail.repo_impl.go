package repoimpl

import (
	"coffee_api/commons"
	verifymail "coffee_api/modules/verify_mail"
	"coffee_api/modules/verify_mail/entity"
	"errors"
	"log"

	"context"

	"gorm.io/gorm"
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
		Email: param.Email,
	}

	result := db.Where(&doc).First(&doc)

	log.Println("this is result: ", result)
	if result.Error != nil || result.RowsAffected == 0 {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			doc.FullName = param.FullName
			doc.SecretCode = param.SecretCode
			if err := db.Create(&doc).Error; err != nil {
				return err
			}
		}

		return result.Error
	}
	return nil
}
