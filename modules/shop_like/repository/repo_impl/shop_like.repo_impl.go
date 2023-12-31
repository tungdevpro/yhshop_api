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
	// shopRepo shop.Repository
}

func NewShopLikeRepoImpl(appCtx commons.AppContext) shoplike.Repository {
	return &shopLikeRepoImpl{
		appCtx: appCtx,
		// shopRepo: shopRepo,
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

	return result, nil
}

func (impl *shopLikeRepoImpl) GetLikedUsers(ctx context.Context, filter *entity.Filter, paging *commons.Paging) ([]commons.SimpleUser, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	shopLikes := []entity.ShopLike{}

	db := impl.appCtx.GetDB()
	db = db.Table(entity.ShopLike{}.TableName())

	if v := filter; v != nil {
		if v.ShopId > 0 {
			db = db.Where("shop_id = ?", v.ShopId)
		}
	}

	db = db.Preload("User")

	if err := db.Find(&shopLikes).Error; err != nil {
		return []commons.SimpleUser{}, err
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return []commons.SimpleUser{}, err
	}

	items := make([]commons.SimpleUser, len(shopLikes))

	for i := range shopLikes {
		shopLikes[i].User.UpdatedAt = nil
		items[i] = *shopLikes[i].User
		items[i].GenerateID()
	}

	return items, nil
}

func (impl *shopLikeRepoImpl) CreateUserLike(ctx context.Context, userId, shopId int) (string, error) {
	// impl.appCtx.L.Lock()
	// defer impl.appCtx.L.Unlock()

	db := impl.appCtx.GetDB()

	data := entity.ShopLike{
		ShopId: shopId,
		UserId: userId,
	}

	if err := db.Table(entity.ShopLike{}.TableName()).Create(&data).Error; err != nil {
		return "", err
	}

	// Ở case này sẽ không tốn time khi call update liked_count, nhưng với những trường hợp khác dùng goroutine sẽ tốt hơn
	// go func() {
	// 	defer commons.Recover(impl.appCtx)
	// 	_ = impl.shopRepo.IncrementLikeCount(ctx, data.ShopId)
	// }()

	return fmt.Sprintf("%d", data.ShopId), nil
}

func (impl *shopLikeRepoImpl) DeleteUserLike(ctx context.Context, userId, shopId int) error {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	db := impl.appCtx.GetDB()

	if err := db.Table(entity.ShopLike{}.TableName()).Where("user_id = ? and shop_id = ?", userId, shopId).Delete(nil).Error; err != nil {
		return err
	}

	// go func() {
	// 	defer commons.Recover(impl.appCtx)
	// 	_ = impl.shopRepo.DecrementLikeCount(ctx, shopId)
	// }()

	return nil
}
