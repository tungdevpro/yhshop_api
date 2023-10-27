package consumers

import (
	"coffee_api/commons"
	authEntity "coffee_api/modules/auth/entity"
	"strconv"

	"coffee_api/modules/verify_mail/entity"
	repoimpl "coffee_api/modules/verify_mail/repository/repo_impl"
	"coffee_api/pubsub"
	"context"
)

func VerifyOTPEmailAfterRegister(appCtx commons.AppContext) consumerJob {
	impl := repoimpl.NewVerifyMail(appCtx)
	return consumerJob{
		Title: "Verify otp email after register",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			doc := message.Data().(*authEntity.OTPRequest)
			param := entity.VerifyMail{
				Email:      doc.Email,
				SecretCode: strconv.Itoa(doc.Otp),
			}

			impl.CheckOTPMail(ctx, param)
			return nil
		},
	}
}
