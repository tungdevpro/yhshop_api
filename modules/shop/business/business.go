package business

import (
	"coffee_api/commons"
	"coffee_api/modules/shop"
	"coffee_api/modules/shop/entity"
	"context"
	"strings"

	"github.com/indrasaputra/hashids"
)

type business struct {
	repo shop.Repository
}

func NewBusiness(repo shop.Repository) shop.Business {
	return &business{repo: repo}

}

func (biz *business) GetListShop(ctx context.Context, filter *entity.Filter, page *commons.Paging) ([]entity.Shop, error) {
	page.Process()
	items, err := biz.repo.GetListShop(ctx, filter, page, "User")
	if err != nil {
		return nil, err
	}

	// ids := make([]int, len(items))
	// for i, e := range items {
	// 	ids[i] = e.Id
	// }

	// mapRespLikes, err := biz.like.GetShopLikes(ctx, ids)
	// if err != nil {
	// 	fmt.Println("cannot get like count")
	// }

	// if v := mapRespLikes; v != nil {
	// 	for i, e := range items {
	// 		items[i].LikedCount = mapRespLikes[e.Id]
	// 	}
	// }

	return items, nil
}
func (biz *business) GetShopById(ctx context.Context, id string) (*entity.Shop, error) {
	xId, err := hashids.DecodeHash([]byte(id))
	if err != nil {
		return nil, err
	}

	item, err := biz.repo.GetShopById(ctx, int(xId))
	if err != nil {
		return nil, err
	}

	item.Mask(false)
	return item, nil
}
func (biz *business) CreateShop(ctx context.Context, dto *entity.CreateShopDTO) (string, error) {
	dto.Name = strings.TrimSpace(dto.Name)
	result, err := biz.repo.CreateShop(ctx, dto)
	if err != nil {
		return "", err
	}

	return result, nil
}
func (biz *business) DeleteShop(ctx context.Context, id string) bool {
	return false
}

func (biz *business) IncrementLikeCount(ctx context.Context, id int) error {
	err := biz.repo.IncrementLikeCount(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (biz *business) DecrementLikeCount(ctx context.Context, id int) error {
	err := biz.repo.DecrementLikeCount(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
