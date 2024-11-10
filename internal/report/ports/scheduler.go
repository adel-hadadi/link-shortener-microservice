package ports

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/report/app"
	"github.com/go-co-op/gocron/v2"
)

const (
	ErrNewScheduler = "error on get new scheduler: %w"
)

func RunScheduler(app app.Application) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(fmt.Errorf(ErrNewScheduler, err))
	}

	startTime := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		59,
		59,
		0,
		time.Local,
	)

	_, err = scheduler.NewJob(
		gocron.DurationJob(time.Hour),
		gocron.NewTask(func() {
			log.Println("Scheduler is running at ", time.Now())

			if err := app.Commands.GenerateReport.Handle(
				context.Background(),
				time.Now(),
			); err != nil {
				log.Fatal(err)
			}
		}),
		gocron.WithStartAt(gocron.WithStartDateTime(startTime)),
	)
	if err != nil {
		log.Println(err)
	}

	scheduler.Start()
}
