package consumers

import (
	"coffee_api/pubsub"
	"context"
	"log"
)

func SendOTPEmailAfterRegister() consumerJob {
	return consumerJob{
		Title: "send otp email after register=======",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			log.Println("heheh::: ", message.Data())

			return nil
		},
	}
}
