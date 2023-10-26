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
	return nil
}
