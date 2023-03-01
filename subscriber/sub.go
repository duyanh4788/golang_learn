package subscriber

import (
	"context"
	"golang_01/common"
	"golang_01/component"
	"golang_01/component/asyncjob"
	"golang_01/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hdl   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx component.AppContext
}

func NewEngine(appContext component.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appContext}
}

func (e *consumerEngine) Start() error {
	e.startSubTopic(common.TopicUserLikeRestaurant, true, RunIncreaseUserLikedRestaurant(e.appCtx))
	e.startSubTopic(common.TopicUserUnLikeRestaurant, true, RunDecreaseUserUnlikedRestaurant(e.appCtx))
	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (e *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := e.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, i := range consumerJobs {
		log.Println("Setup consumer for", i.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("run job for", job.Title, ", Value:", message.Data())
			return job.Hdl(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdlArr[i] = asyncjob.NewJob(getJobHandler(&consumerJobs[i], msg))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
