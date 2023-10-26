package consumers

import (
	"coffee_api/commons"
	asyncjob "coffee_api/commons/async_job"
	"coffee_api/pubsub"
	"context"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx commons.AppContext
}

func NewCusumnerEngine(appCtx commons.AppContext) *consumerEngine {
	return &consumerEngine{
		appCtx: appCtx,
	}
}

func (engine *consumerEngine) Start() error {
	err := engine.startSubTopic(commons.ChanVerifyMailCreated, false, SendOTPEmailAfterRegister(engine.appCtx))
	if err != nil {
		return err
	}
	return nil
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Channel, isParallel bool, hdls ...consumerJob) error {

	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range hdls {
		log.Println("Setup consumer for:", item.Title)
	}

	getHld := func(job *consumerJob, message *pubsub.Message) func(ctx context.Context) error {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(hdls))

			for i := range hdls {
				jobHdlArr[i] = asyncjob.NewJob(getHld(&hdls[i], msg))
			}

			group := asyncjob.NewGroup(isParallel, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
