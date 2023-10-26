package subscriber

import (
	"coffee_api/commons"
	implShop "coffee_api/modules/shop/repository/repo_impl"
	"coffee_api/modules/shop_like/entity"

	"context"
)

func IncreaseLikeCountAfterUserLike(appCtx commons.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, commons.TopicUserLikeShop)
	impl := implShop.NewShopRepoImpl(appCtx)

	go func() {
		for {
			msg := <-c
			data := msg.Data().(*entity.ShopLike)
			_ = impl.IncrementLikeCount(ctx, data.ShopId)
		}
	}()

}
