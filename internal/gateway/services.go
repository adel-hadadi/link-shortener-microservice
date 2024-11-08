package main

import "context"

type LinkService interface {
	CreateLink(ctx context.Context, originalURL string) (string, error)
	GetLink(ctx context.Context, shortURL string) (string, error)
}
