package ports

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/adel-hadadi/link-shotener/internal/report/app"
	"github.com/go-co-op/gocron/v2"
)

const (
	ErrNewScheduler = "error on get new scheduler: %w"
	reportSheet     = "Clicks Report"
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

	now := time.Now().Format("2006-01-02")
	filePath := fmt.Sprintf("./storage/clicks-report-%s.xlsx", now)

	f, err := getReportFile(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err := f.SaveAs(filePath); err != nil {
			log.Fatal(err)
		}
		f.Close()
	}()

	rows, err := f.GetRows(reportSheet)
	if err != nil {
		log.Println(err)
		return
	}

	currentHour := time.Now().Hour()

	for _, click := range clicks {
		if index, exists := checkShortURLExists(click.ShortURL, rows); exists {
			cell, _ := excelize.CoordinatesToCellName(currentHour+2, index+1)
			f.SetCellValue(reportSheet, cell, click.Count)
			continue
		}

		cell, _ := excelize.CoordinatesToCellName(1, len(rows)+1)
		f.SetCellValue(reportSheet, cell, click.ShortURL)

		countCell, _ := excelize.CoordinatesToCellName(currentHour+2, len(rows)+1)
		f.SetCellValue(reportSheet, countCell, click.Count)
	}
}

func getReportFile(filePath string) (*excelize.File, error) {
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f := excelize.NewFile()

			sheetIndex, err := f.NewSheet(reportSheet)
			if err != nil {
				log.Println(err)
			}

			headers := []string{"Short URL"}
			for hour := 0; hour < 24; hour++ {
				headers = append(headers, fmt.Sprintf("%02d:00", hour))
			}

			for i, header := range headers {
				cell, _ := excelize.CoordinatesToCellName(i+1, 1)
				f.SetCellValue(reportSheet, cell, header)
			}

			f.SetActiveSheet(sheetIndex)

			return f, nil
		}
	}

	return excelize.OpenFile(filePath)
}

func checkShortURLExists(shortURL string, rows [][]string) (int, bool) {
	for i, row := range rows {
		if row[0] == shortURL {
			return i, true
		}
	}

	return 0, false
}
