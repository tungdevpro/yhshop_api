package subscriber

import (
	"coffee_api/commons"
	"context"
)

func Setup(appCtx commons.AppContext) {
	IncreaseLikeCountAfterUserLike(appCtx, context.Background())
}
