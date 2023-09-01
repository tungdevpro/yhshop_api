package repoimpl

import (
	"coffee_api/modules/auth/entity"
	"coffee_api/modules/auth/repository"
	"context"

	"gorm.io/gorm"
)

type authRepoImpl struct {
	db *gorm.DB
}


func NewAuthRepoImpl(db *gorm.DB) repository.Repository {
	return &authRepoImpl{
		db: db,
	}
}


func (repo *authRepoImpl) Register(ctx context.Context, request entity.RegisterRequest) {}
func (repo *authRepoImpl) Login(ctx context.Context, request entity.LoginRequest) {}
