package subscriber

import (
	"coffee_api/commons"
	"context"
	"log"
)

func IncreaseLikeCountAfterUserLike(appCtx commons.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, commons.TopicUserLikeShop)
	go func() {
		for {
			msg := <-c
			log.Println("New event published:", msg.Data())
		}
	}()

}
