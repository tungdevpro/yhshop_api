package business

import (
	"coffee_api/modules/user"
	"context"
)

type business struct {
	repo user.Repository
}

func NewBusiness(repo user.Repository) user.Business {
	return &business{repo: repo}
}

func (biz *business) GetProfile(ctx context.Context)    {}
func (biz *business) DelUser(ctx context.Context)       {}
func (biz *business) UpdateProfile(ctx context.Context) {}
