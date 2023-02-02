package main

import (
	"context"
	"golang_01/component/asyncjob"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("Im job 1")
		return nil
		//return errors.New("something went wrong job 1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("Im job 2")

		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		log.Println("Im job 3")
		return nil
	})

	group := asyncjob.NewGroup(true, job1, job2, job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}

	//if err := job1.Execute(context.Background()); err != nil {
	//	log.Println("Start job 1", err)
	//
	//	for {
	//		if err := job1.Retry(context.Background()); err == nil {
	//			break
	//		}
	//
	//		if job1.State() == asyncjob.StateRetryFailed {
	//			break
	//		}
	//	}
	//}
}
