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

func (impl *shopRepoImpl) GetListShop(ctx context.Context, filter *entity.Filter, page *commons.Paging) ([]entity.Shop, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	var items []entity.Shop
	db := impl.appCtx.GetDB()

	db = db.Table(entity.Shop{}.TableName()).Where("status = 1")
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Limit(page.Limit).Offset((page.Page - 1) * page.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	for i := range items {
		items[i].Mask(false)
	}

	return items, nil
}
func (impl *shopRepoImpl) GetShopById(ctx context.Context) (entity.Shop, error) {
	return entity.Shop{}, nil
}
func (impl *shopRepoImpl) CreateShop(ctx context.Context, dto *entity.CreateShopDTO) (string, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

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

	return "", nil
}
func (impl *shopRepoImpl) DeleteShop(ctx context.Context) {}
