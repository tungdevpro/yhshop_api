package repository

import (
	"coffee_api/modules/auth/entity"
	"context"
)


type Repository interface {
	Login(ctx context.Context, req entity.LoginRequest)
	Register(ctx context.Context, req entity.RegisterRequest)
}
