package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/auth/entity"
	"coffee_api/pubsub"
	"context"
)

func (impl *authRepoImpl) VerifyOTP(ctx context.Context, param *entity.OTPRequest) (bool, error) {
	err := impl.appCtx.GetPubsub().Publish(ctx, commons.ChanVerifyMailCreated, pubsub.NewMessage(param))
	if err != nil {
		return false, err
	}

	return true, nil
}
