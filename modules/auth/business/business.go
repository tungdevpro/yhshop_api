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

func NewBusiness(r auth.Repository) auth.Business {
	return &business{
		repository: r,
	}
}

func (biz *business) Register(ctx context.Context, req *entity.RegisterDTO) error {
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

	return nil
}

func (biz *business) Login(ctx context.Context, req *entity.LoginDTO) error {
	if err := req.Validate(); err != nil {
		return err
	}

	if err := biz.repository.Login(ctx, req); err != nil {
		return err
	}

	fmt.Println("this login: ", req)

	return nil
}
