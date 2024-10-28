package service

import (
	"context"
	"crypto/sha1"
	"encoding/base64"

	apperr "github.com/adel-hadadi/link-shotener/internal/common/errors"
)

type LinkService struct {
	repo linkRepository
}

type CreateLink struct {
	OriginalLink string
}

type linkRepository interface {
	Create(ctx context.Context, originalLink, shortLink string) error
	GetByShortLink(ctx context.Context, link string) (Link, error)
}

func NewLinkService(repo linkRepository) LinkService {
	return LinkService{
		repo: repo,
	}
}

func (s LinkService) Create(ctx context.Context, req CreateLink) error {
	shortURL := generateShortCode(req.OriginalLink)

	err := s.repo.Create(ctx, req.OriginalLink, shortURL)
	if err != nil {
		if apperr.IsSQLDuplicateEntry(err) {
			return apperr.NewConflictError(err.Error(), "link-already-exists")
		}

		return err
	}

	return nil
}

func generateShortCode(url string) string {
	hash := sha1.Sum([]byte(url))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:6] // Limit to 6 characters
	return shortURL
}

func (s LinkService) Get(ctx context.Context, shortLink string) (Link, error) {
	link, err := s.repo.GetByShortLink(ctx, shortLink)
	if err != nil {
		return Link{}, apperr.NewSlugError(err.Error(), "not-found")
	}

	return link, nil
}
