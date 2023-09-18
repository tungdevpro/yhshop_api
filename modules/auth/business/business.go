package business

import (
	"coffee_api/modules/auth"
	"coffee_api/modules/auth/entity"
	"context"

	"github.com/indrasaputra/hashids"
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

func (biz *business) Register(ctx context.Context, req *entity.RegisterDTO) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	req.Password = string(passwordHash[:])
	result, err := biz.repository.Register(ctx, req)
	if err != nil {
		return "", err
	}

	return result, err
}

func (biz *business) Login(ctx context.Context, req *entity.LoginDTO) (*entity.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	result, err := biz.repository.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	xId, _ := hashids.EncodeID(hashids.ID(result.Id))

	result.Uid = string(xId)

	return result, nil
}
