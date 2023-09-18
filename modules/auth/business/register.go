package business

import (
	"coffee_api/modules/auth/entity"
	"context"

	"golang.org/x/crypto/bcrypt"
)

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
