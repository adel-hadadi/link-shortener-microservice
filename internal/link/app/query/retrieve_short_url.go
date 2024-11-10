package query

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/link/domain/link"
)

type RetrieveShortURLHandler struct {
	repo link.Repository
}

func NewRetrieveShortURLHandler(repo link.Repository) RetrieveShortURLHandler {
	return RetrieveShortURLHandler{repo: repo}
}

func (h RetrieveShortURLHandler) Handle(ctx context.Context, originalURL string) (string, error) {
	return h.repo.GetByOriginalURL(ctx, originalURL)
}
