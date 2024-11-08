package service

import (
	"context"
	"time"
)

type ReportService struct {
	repo hourlyClickRepository
}

type hourlyClickRepository interface {
	Create(ctx context.Context, shortURL string, clickedAt time.Time) error
	GetLastHourClicks(ctx context.Context) ([]ClickCount, error)
}

func NewReportService(repo hourlyClickRepository) ReportService {
	return ReportService{
		repo: repo,
	}
}

type LinkClickedReq struct {
	ClickedAt time.Time
	Link      string
}

func (s ReportService) LinkClicked(ctx context.Context, req LinkClickedReq) error {
	return s.repo.Create(ctx, req.Link, req.ClickedAt)
}

func (s ReportService) GetLastHourClicks(ctx context.Context) ([]ClickCount, error) {
	clicks, err := s.repo.GetLastHourClicks(ctx)
	if err != nil {
		return nil, err
	}

	return clicks, nil
}
