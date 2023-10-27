package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/auth/entity"
	"coffee_api/pubsub"
	"context"
)

func (impl *authRepoImpl) VerifyOTP(ctx context.Context, param *entity.OTPRequest) (bool, error) {
	// db := impl.appCtx.GetDB()

	// doc := entity.OTPRequest{
	// 	Email: param.Email,
	// }

	// result := db.Where(&doc).Find(&doc)
	// fmt.Println("result....", result)
	// if result.Error == nil || result.RowsAffected != 0 {
	// 	return false, result.Error
	// }

	// if doc.Otp != param.Otp {
	// 	return false, entity.ErrOTPNotEqual
	// }

	impl.appCtx.GetPubsub().Publish(ctx, commons.ChanVerifyMailCreated, pubsub.NewMessage(param))

	return false, nil
}
