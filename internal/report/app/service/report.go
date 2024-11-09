package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/xuri/excelize/v2"
)

const reportSheet = "Report Sheet"

type ReportService struct {
	repo    hourlyClickRepository
	storage reportStorage
}

type hourlyClickRepository interface {
	Create(ctx context.Context, shortURL string, clickedAt time.Time) error
	GetLastHourClicks(ctx context.Context) ([]ClickCount, error)
}

type reportStorage interface {
	GetObject(ctx context.Context, filePath string) (*bytes.Buffer, error)
	PutObject(ctx context.Context, content *bytes.Buffer, filePath string) error
}

func NewReportService(repo hourlyClickRepository, storage reportStorage) ReportService {
	return ReportService{
		repo:    repo,
		storage: storage,
	}
}

type LinkClickedReq struct {
	ClickedAt time.Time
	Link      string
}

func (s ReportService) LinkClicked(ctx context.Context, req LinkClickedReq) error {
	return s.repo.Create(ctx, req.Link, req.ClickedAt)
}

func (s ReportService) GenerateClicksReport(ctx context.Context, currentTime time.Time) error {
	clicks, err := s.repo.GetLastHourClicks(ctx)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("clicks-report-%s.xlsx", currentTime.Format("2006-01-02"))

	buf, err := s.storage.GetObject(ctx, filePath)
	if err != nil {
		minioErr := minio.ToErrorResponse(err)
		if minioErr.Code == "NoSuchKey" {
			buf, _ = s.generateReportFile()
		} else {
			return fmt.Errorf("error on get object from storage: %w", err)
		}
	}

	f, err := excelize.OpenReader(buf)
	if err != nil {
		return err
	}
	defer func() {
		var output bytes.Buffer
		if err := f.Write(&output); err != nil {
			log.Fatal(err)
		}

		if err := s.storage.PutObject(ctx, &output, filePath); err != nil {
			log.Fatal(err)
		}
		f.Close()
	}()

	rows, err := f.GetRows(reportSheet)
	if err != nil {
		return err
	}

	currentHour := time.Now().Hour()

	for _, click := range clicks {
		if index, exists := s.checkShortURLExists(click.ShortURL, rows); exists {
			cell, _ := excelize.CoordinatesToCellName(currentHour+2, index+1)
			f.SetCellValue(reportSheet, cell, click.Count)
			continue
		}

		cell, _ := excelize.CoordinatesToCellName(1, len(rows)+1)
		f.SetCellValue(reportSheet, cell, click.ShortURL)

		countCell, _ := excelize.CoordinatesToCellName(currentHour+2, len(rows)+1)
		f.SetCellValue(reportSheet, countCell, click.Count)
	}

	return nil
}

func (s ReportService) DownloadReport(ctx context.Context, filePath string) (*bytes.Buffer, error) {
	content, err := s.storage.GetObject(ctx, filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s ReportService) generateReportFile() (*bytes.Buffer, error) {
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

	var buf bytes.Buffer
	err = f.Write(&buf)

	return &buf, err
}

func (s ReportService) checkShortURLExists(shortURL string, rows [][]string) (int, bool) {
	for i, row := range rows {
		if row[0] == shortURL {
			return i, true
		}
	}

	return 0, false
}
