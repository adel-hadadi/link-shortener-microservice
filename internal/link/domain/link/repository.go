package link

import "context"

type Repository interface {
	Create(ctx context.Context, originalURL string, shortURL string) error
	GetByOriginalURL(ctx context.Context, originalURL string) (string, error)
	GetByShortURL(ctx context.Context, shortURL string) (string, error)
}
