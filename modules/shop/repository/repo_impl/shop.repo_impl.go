package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"context"
	"fmt"
)

type shopRepoImpl struct {
	appCtx commons.AppContext
}

func NewShopRepoImpl(appCtx commons.AppContext) shop.Repository {
	return &shopRepoImpl{
		appCtx: appCtx,
	}
}

func (impl *shopRepoImpl) GetListShop(ctx context.Context, filter *entity.Filter, page *commons.Paging) ([]entity.Shop, error) {
	return []entity.Shop{}, nil
}
func (impl *shopRepoImpl) GetShopById(ctx context.Context) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (impl *shopRepoImpl) CreateShop(ctx context.Context, dto *entity.CreateShopDTO) (string, error) {
	db := impl.appCtx.GetDB()

	shop := entity.Shop{
		Name:     dto.Name,
		CityId:   1,
		OwnerId:  2,
		IsVerify: 1,
	}

	if err := db.Create(&shop).Error; err != nil {
		return "", nil
	}

	shop.Mask(false)

	fmt.Println("shop>>> ", shop)

	return "", nil
}
func (impl *shopRepoImpl) DeleteShop(ctx context.Context) {}
