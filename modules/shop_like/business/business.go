package business

import (
	shoplike "coffee_api/modules/shop_like"
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
