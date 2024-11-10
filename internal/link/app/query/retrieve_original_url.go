package query

import (
	"context"

	"github.com/adel-hadadi/link-shotener/internal/link/domain/link"
)

type RetrieveOriginalURLHandler struct {
	repo link.Repository
}

func NewRetrieveOriginalURLHandler(repo link.Repository) RetrieveOriginalURLHandler {
	return RetrieveOriginalURLHandler{repo: repo}
}

func (h RetrieveOriginalURLHandler) Handle(ctx context.Context, shortURL string) (string, error) {
	return h.repo.GetByShortURL(ctx, shortURL)
}
