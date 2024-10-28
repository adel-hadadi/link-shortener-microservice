package service

import (
	"context"
)

type ClickService struct {
	repo clickRepository
}

type clickRepository interface {
	Create(ctx context.Context, linkID int, IPAddress string) error
}

func NewClickService(repo clickRepository) ClickService {
	return ClickService{
		repo: repo,
	}
}

type CreateClick struct {
	LinkID    int
	IPAddress string
}

func (s ClickService) Create(ctx context.Context, params CreateClick) error {
	return s.repo.Create(ctx, params.LinkID, params.IPAddress)
}
