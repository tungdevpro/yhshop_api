package consumers

import (
	"coffee_api/commons"

	authEntity "coffee_api/modules/auth/entity"
	repoimpl "coffee_api/modules/user/repository/repo_impl"
	"coffee_api/pubsub"
	"context"
)

func ChangeIsVerifyEmailInUserModel(appCtx commons.AppContext) consumerJob {
	impl := repoimpl.NewUserRepoImpl(appCtx)
	return consumerJob{
		Title: "Verify otp email after register",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			doc := message.Data().(*authEntity.OTPRequest)

			_ = impl.ChangeVerifyEmail(ctx, doc.Email)
			return nil
		},
	}
}
