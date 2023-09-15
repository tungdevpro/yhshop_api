package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"context"
)

type shopRepoImpl struct {
	appCtx commons.AppContext
}

func NewShopRepoImpl(appCtx commons.AppContext) shop.Repository {
	return &shopRepoImpl{
		appCtx: appCtx,
	}
}

func (impl *shopRepoImpl) GetListShop(ctx context.Context) {}
func (impl *shopRepoImpl) GetShopById(ctx context.Context) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (impl *shopRepoImpl) CreateShop(ctx context.Context, dto *entity.CreateShopDTO) (string, error) {
	return "", nil
}
func (impl *shopRepoImpl) DeleteShop(ctx context.Context) {}
