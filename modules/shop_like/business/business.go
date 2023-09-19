package business

import (
	"coffee_api/commons"
	shoplike "coffee_api/modules/shop_like"
	"coffee_api/modules/shop_like/entity"
	"context"
	"fmt"
)

type business struct {
	repo shoplike.Repository
}

func NewBusiness(repo shoplike.Repository) shoplike.Business {
	return &business{
		repo: repo,
	}
}

func (biz *business) GetShopLikes(ctx context.Context, ids []int) (map[int]int, error) {
	item, err := biz.repo.GetShopLikes(ctx, ids)
	if err != nil {
		fmt.Println("has error")
	}

	return item, nil
}

func (biz *business) GetLikedUsers(ctx context.Context, filter *entity.Filter) ([]commons.SimpleUser, error) {

	return nil, nil
}
