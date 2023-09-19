package jwtexplore

import (
	"coffee_api/modules/user/entity"
	"context"
)

type Repository interface {
	FindUser(context.Context, int) (*entity.User, error)
}

type Business interface {
	FindUser(context.Context, int) (*entity.User, error)
}
