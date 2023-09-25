package subscriber

import (
	"coffee_api/commons"
	"context"
	"fmt"
)

func IncreaseLikeCountAfterUserLike(appCtx commons.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, commons.TopicUserLikeShop)

	go func() {
		for {
			msg := <-c
			fmt.Println("msg: >> ", msg)
		}
	}()
}
