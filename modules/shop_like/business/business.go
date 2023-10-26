package business

import (
	"coffee_api/commons"
	asyncjob "coffee_api/commons/async_job"
	"coffee_api/modules/shop"
	shoplike "coffee_api/modules/shop_like"
	"coffee_api/modules/shop_like/entity"
	"context"
	"fmt"
)

type business struct {
	repo    shoplike.Repository
	bizShop shop.Business
	// pubsub  pubsub.Pubsub
}

func NewBusiness(repo shoplike.Repository, bizShop shop.Business) shoplike.Business {
	return &business{
		repo:    repo,
		bizShop: bizShop,
		// pubsub:  pubsub,
	}
}

func (biz *business) GetShopLikes(ctx context.Context, ids []int) (map[int]int, error) {
	item, err := biz.repo.GetShopLikes(ctx, ids)
	if err != nil {
		fmt.Println("has error")
	}

	return item, nil
}

func (biz *business) GetLikedUsers(ctx context.Context, filter *entity.Filter, paging *commons.Paging) ([]commons.SimpleUser, error) {
	items, err := biz.repo.GetLikedUsers(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (biz *business) CreateUserLike(ctx context.Context, userId, shopId int) (string, error) {
	id, err := biz.repo.CreateUserLike(ctx, userId, shopId)
	if err != nil {
		return "", err
	}

	// data := entity.Filter{
	// 	ShopId: shopId,
	// 	UserId: userId,
	// }

	// _ = biz.pubsub.Publish(ctx, commons.TopicUserLikeShop, pubsub.NewMessage(data))
	// job := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.bizShop.IncrementLikeCount(ctx, shopId)
	// })

	// _ = asyncjob.NewGroup(true, job).Run(ctx)
	return id, nil
}

func (biz *business) DeleteUserLike(ctx context.Context, userId, shopId int) error {
	if err := biz.repo.DeleteUserLike(ctx, userId, shopId); err != nil {
		return err
	}

	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.bizShop.DecrementLikeCount(ctx, shopId)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
