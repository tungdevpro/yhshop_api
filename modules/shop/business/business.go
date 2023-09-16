package business

import (
	"coffee_api/commons"
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"context"
	"strings"
)

type business struct {
	repo shop.Repository
}

func NewBusiness(repo shop.Repository) shop.Business {
	return &business{repo: repo}
}

func (biz *business) GetListShop(ctx context.Context, filter *entity.Filter, page *commons.Paging) ([]entity.Shop, error) {

	// items, err := biz.repo.GetListShop(ctx, filter)

	return []entity.Shop{}, nil
}
func (biz *business) GetShopById(ctx context.Context) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (biz *business) CreateShop(ctx context.Context, dto *entity.CreateShopDTO) (string, error) {
	dto.Name = strings.TrimSpace(dto.Name)
	result, err := biz.repo.CreateShop(ctx, dto)
	if err != nil {
		return "", err
	}

	return result, nil
}
func (biz *business) DeleteShop(ctx context.Context) {}
