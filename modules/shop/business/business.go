package business

import (
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"context"
)

type business struct {
	repo shop.Repository
}

func NewBusiness(repo shop.Repository) shop.Business {
	return &business{repo: repo}
}

func (biz *business) GetListShop(ctx context.Context) {}
func (biz *business) GetShopById(ctx context.Context) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (biz *business) CreateShop(ctx context.Context) {}
func (biz *business) DeleteShop(ctx context.Context) {}
