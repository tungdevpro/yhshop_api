package consumers

import (
	"coffee_api/commons"
	"coffee_api/modules/verify_mail/entity"
	repoimpl "coffee_api/modules/verify_mail/repository/repo_impl"
	"coffee_api/pubsub"
	"context"
	"encoding/json"
)

func SendOTPEmailAfterRegister(appCtx commons.AppContext) consumerJob {
	impl := repoimpl.NewVerifyMail(appCtx)
	return consumerJob{
		Title: "Send otp email after register",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			dict := make(map[string]interface{})

			for k, v := range message.Data().(map[string]interface{}) {
				dict[k] = v
			}

			jsonbody, err := json.Marshal(dict)
			if err != nil {
				return err
			}
			param := entity.VerifyMail{}
			if err := json.Unmarshal(jsonbody, &param); err != nil {
				return err
			}

			_ = impl.CreateMail(ctx, param)

			return nil
		},
	}
}
