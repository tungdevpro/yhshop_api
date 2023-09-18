package business

import (
	shoplike "coffee_api/modules/shop_like"
	"context"
)

type business struct {
}

func NewBusiness() shoplike.Business {
	return &business{}
}

func (biz *business) GetShopLikes(ctx context.Context, ids []int) (map[int]int, error) {
	return nil, nil
}
