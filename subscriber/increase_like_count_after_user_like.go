package subscriber

import (
	"coffee_api/commons"
	implShop "coffee_api/modules/shop/repository/repo_impl"
	"coffee_api/modules/shop_like/entity"
	"fmt"

	"context"
	"log"
)

func IncreaseLikeCountAfterUserLike(appCtx commons.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, commons.TopicUserLikeShop)
	impl := implShop.NewShopRepoImpl(appCtx)

	go func() {
		for {
			msg := <-c
			data := msg.Data().(*entity.ShopLike)
			err := impl.IncrementLikeCount(ctx, data.ShopId)
			fmt.Println("err:>> ", err)
			// storage.CreateUserLike(ctx, data.UserId, data.ShopId)
			log.Println("Pushed:", msg.Data())
		}
	}()

}
