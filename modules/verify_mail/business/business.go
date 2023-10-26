package business

import (
	verifymail "coffee_api/modules/verify_mail"
	"coffee_api/modules/verify_mail/entity"
	"context"
)

type business struct {
	repo verifymail.Repository
}

func NewBusiness(repo verifymail.Repository) verifymail.Business {
	return &business{
		repo: repo,
	}
}

func (biz *business) CreateMail(context context.Context, param entity.VerifyMail) error {
	return nil
}
