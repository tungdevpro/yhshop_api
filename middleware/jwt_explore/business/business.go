package business

import (
	"coffee_api/commons"
	jwtexplore "coffee_api/middleware/jwt_explore"
	"context"
)

type business struct {
	repo jwtexplore.Repository
}

func NewBusiness(repo jwtexplore.Repository) jwtexplore.Business {
	return &business{
		repo: repo,
	}
}

func (biz *business) FindUser(context.Context, int) (*commons.SimpleUser, error) {
	return nil, nil
}
