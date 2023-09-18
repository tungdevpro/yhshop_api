package repoimpl

import (
	"coffee_api/commons"
	shoplike "coffee_api/modules/shop_like"
	"coffee_api/modules/shop_like/entity"
	"context"
	"fmt"
)

type shopLikeRepoImpl struct {
	appCtx commons.AppContext
}

func NewShopLikeRepoImpl(appCtx commons.AppContext) shoplike.Business {
	return &shopLikeRepoImpl{
		appCtx: appCtx,
	}
}

func (impl *shopLikeRepoImpl) GetShopLikes(ctx context.Context, ids []int) (map[int]int, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	result := make(map[int]int)

	type sqlData struct {
		ShopId    int `json:"shop_id" gorm:"column:shop_id;"`
		LikeCount int `json:"count" gorm:"column:count;"`
	}

	var listLike []sqlData
	db := impl.appCtx.GetDB()

	if err := db.Table(entity.ShopLike{}.TableName()).Select("shop_id, count(shop_id) as count").Where("shop_id in (?)", ids).Group("shop_id").Find(&listLike).Error; err != nil {
		return nil, err
	}

	for _, item := range listLike {
		result[item.ShopId] = item.LikeCount
	}

	fmt.Println("impl:::: ", result)

	return result, nil
}
