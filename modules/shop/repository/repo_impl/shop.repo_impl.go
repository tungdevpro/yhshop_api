package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"context"
	"fmt"

	"github.com/indrasaputra/hashids"
	"gorm.io/gorm"
)

type shopRepoImpl struct {
	appCtx commons.AppContext
}

func NewShopRepoImpl(appCtx commons.AppContext) shop.Repository {
	return &shopRepoImpl{
		appCtx: appCtx,
	}
}

func (impl *shopRepoImpl) GetListShop(ctx context.Context, filter *entity.Filter, paging *commons.Paging, args ...string) ([]entity.Shop, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	var items []entity.Shop
	db := impl.appCtx.GetDB()

	db = db.Table(entity.Shop{}.TableName()).Where("status = 1")
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range args {
		db = db.Preload(args[i])
	}

	if paging.FakeCursor != "" {
		if curId, err := hashids.DecodeHash([]byte(paging.FakeCursor)); err == nil {
			db = db.Where("id < ?", curId)
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&items).Error; err != nil {
		return nil, err
	}

	for i := range items {
		items[i].Mask(false)

		if i == len(items)-1 {
			paging.NextCursor = items[i].Uid.EncodeString()
		}
	}

	return items, nil
}
func (impl *shopRepoImpl) GetShopById(ctx context.Context, id int) (*entity.Shop, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	item := entity.Shop{}

	db := impl.appCtx.GetDB()
	if err := db.First(&item, id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
func (impl *shopRepoImpl) CreateShop(ctx context.Context, dto *entity.CreateShopDTO) (string, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	db := impl.appCtx.GetDB()

	shop := entity.Shop{
		Name:     dto.Name,
		CityId:   1,
		UserId:   dto.OwnerId,
		IsVerify: 0,
	}

	if err := db.Create(&shop).Error; err != nil {
		return "", nil
	}

	shop.Mask(false)

	return "", nil
}
func (impl *shopRepoImpl) DeleteShop(ctx context.Context, id string) bool {
	return true
}

func (impl *shopRepoImpl) IncrementLikeCount(ctx context.Context, id int) error {
	db := impl.appCtx.GetDB()

	fmt.Println("what your name....")

	if err := db.Table(entity.Shop{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func (impl *shopRepoImpl) DecrementLikeCount(ctx context.Context, shopId int) error {
	db := impl.appCtx.GetDB()

	if err := db.Table(entity.Shop{}.TableName()).Where("id = ?", shopId).Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return err
	}

	return nil
}
