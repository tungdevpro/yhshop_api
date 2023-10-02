package subscriber

import (
	"coffee_api/commons"
	"context"
)

func Setup(appCtx commons.AppContext) {
	IncreaseLikeCountAfterUserLike(appCtx, context.Background())
}

type engineSubscriber struct {
	appCtx commons.AppContext
}

func NewEngineSubscriber(appCtx commons.AppContext) *engineSubscriber {
	return &engineSubscriber{appCtx: appCtx}
}

func (e *engineSubscriber) StartTopic() {}
