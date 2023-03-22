package main

import (
	"context"
	"errors"
	"food-delivery/component/asyncjob"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)

		log.Println("Job 1")
		//return nil
		return errors.New("Occur error when run job 1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)

		log.Println("Job 2")
		//return nil
		return errors.New("Occur error when run job 2")
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)

		log.Println("Job 3")
		//return nil
		return errors.New("Occur error when run job 3")
	})

	group := asyncjob.NewGroup(false, job1, job2, job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
	//job1.SetRetryDuration([]time.Duration{time.Second * 3})
	//
	//if err := job1.Execute(context.Background()); err != nil {
	//	log.Println(job1.State(), err)
	//
	//	for {
	//		if err := job1.Retry(context.Background()); err != nil {
	//			log.Println(err)
	//		}
	//
	//		if job1.State() == asyncjob.StateRetryFailed || job1.State() == asyncjob.StateCompleted {
	//			log.Println(job1.State())
	//			break
	//		}
	//	}
	//}

}
