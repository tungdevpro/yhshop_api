package business

import (
	"coffee_api/modules/auth/repository"
	"context"
)

type IBusiness interface {
	Register(context.Context) error
	Login(context.Context) error
}

type business struct {
	repository repository.Repository
}

func NewBusiness(r repository.Repository) *business {
	return &business{
		repository: r,
	}
}

func (biz *business) Register(ctx context.Context) error {
	return nil
}

func (biz *business) Login(ctx context.Context) error {
	return nil
}
