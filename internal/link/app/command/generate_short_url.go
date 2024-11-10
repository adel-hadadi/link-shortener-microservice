package command

import (
	"context"
	"crypto/sha1"
	"encoding/base64"

	apperr "github.com/adel-hadadi/link-shotener/internal/common/errors"
	"github.com/adel-hadadi/link-shotener/internal/link/domain/link"
)

type GenerateShortURL struct {
	OriginalURL string
}

type GenerateShortURLHandler struct {
	repo link.Repository
}

func NewGenerateShortURLHandler(repo link.Repository) GenerateShortURLHandler {
	return GenerateShortURLHandler{repo: repo}
}

func (h GenerateShortURLHandler) Handle(ctx context.Context, cmd GenerateShortURL) error {
	_, err := h.repo.GetByOriginalURL(ctx, cmd.OriginalURL)
	if err != nil {
		if !apperr.IsSQLNoRows(err) {
			return apperr.NewSlugError(err.Error(), "internal-server-error")
		}

		shortURL := generateShortCode(cmd.OriginalURL)

		return h.repo.Create(ctx, cmd.OriginalURL, shortURL)
	}

	return nil
}

func generateShortCode(url string) string {
	hash := sha1.Sum([]byte(url))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:7]
	return shortURL
}
