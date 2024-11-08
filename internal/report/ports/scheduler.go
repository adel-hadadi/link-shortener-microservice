package ports

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"

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
			generateExcelReport(app)
		}),
		gocron.WithStartAt(gocron.WithStartDateTime(startTime)),
	)
	if err != nil {
		log.Println(err)
	}

	scheduler.Start()
}

func generateExcelReport(app app.Application) {
	log.Println("Scheduler is running at ", time.Now())

	clicks, err := app.Services.ReportService.GetLastHourClicks(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("clicks count => ", len(clicks))

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	sheet := "Clicks Report"
	index, err := f.NewSheet(sheet)
	if err != nil {
		log.Println(err)
	}

	headers := []string{"Short URL"}
	for hour := 0; hour < 24; hour++ {
		headers = append(headers, fmt.Sprintf("%02d:00", hour))
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	for i, click := range clicks {
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		f.SetCellValue(sheet, cell, click.ShortURL)

		countCell, _ := excelize.CoordinatesToCellName(time.Now().Hour()+2, i+2)

		f.SetCellValue(sheet, countCell, click.Count)
	}

	f.SetActiveSheet(index)

	if err := f.SaveAs("storage/report.xlsx"); err != nil {
		log.Fatal(err)
	}
}
