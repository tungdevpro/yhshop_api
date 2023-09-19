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

func (biz *business) FindUser(ctx context.Context, id int) (*commons.SimpleUser, error) {
	result, err := biz.repo.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
