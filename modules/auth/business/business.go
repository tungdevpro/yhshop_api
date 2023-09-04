package business

import (
	"coffee_api/modules/auth"
	"coffee_api/modules/auth/entity"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type business struct {
	repository auth.Repository
}

func NewBusiness(r auth.Repository) *business {
	return &business{
		repository: r,
	}
}

func (biz *business) Register(ctx context.Context, req *entity.RegisterRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(passwordHash[:])

	if err := biz.repository.Register(ctx, req); err != nil {
		return err
	}

	fmt.Println("req....", req)

	return nil
}

func (biz *business) Login(ctx context.Context, req *entity.LoginRequest) error {
	return nil
}
