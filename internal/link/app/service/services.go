package service

import (
	"context"
)

type reportService interface {
	LinkClicked(ctx context.Context, shortURL string) error
}
