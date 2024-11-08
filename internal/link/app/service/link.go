package service

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"time"

	apperr "github.com/adel-hadadi/link-shotener/internal/common/errors"
	"github.com/adel-hadadi/link-shotener/internal/link/adapters"
)

type LinkService struct {
	repo      linkRepository
	reportSvc adapters.ReportGrpc
}

type CreateLink struct {
	OriginalLink string
}

type linkRepository interface {
	Create(ctx context.Context, originalLink, shortLink string) error
	GetByShortLink(ctx context.Context, link string) (Link, error)
	GetByOriginalLink(ctx context.Context, link string) (Link, error)
}

func NewLinkService(repo linkRepository, reportSvc adapters.ReportGrpc) LinkService {
	return LinkService{
		repo:      repo,
		reportSvc: reportSvc,
	}
}

func (s LinkService) Create(ctx context.Context, req CreateLink) (string, error) {
	link, err := s.repo.GetByOriginalLink(ctx, req.OriginalLink)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			shortURL := generateShortCode(req.OriginalLink)

			err := s.repo.Create(ctx, req.OriginalLink, shortURL)
			if err != nil {
				if apperr.IsSQLDuplicateEntry(err) {
					return "", apperr.NewConflictError(err.Error(), "link-already-exists")
				}

				return "", err
			}

			return shortURL, nil
		}

		return "", apperr.NewSlugError(err.Error(), "internal-server-error")
	}

	return link.ShortLink, nil
}

func generateShortCode(url string) string {
	hash := sha1.Sum([]byte(url))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:7]
	return shortURL
}

func (s LinkService) Get(ctx context.Context, shortLink string) (Link, error) {
	link, err := s.repo.GetByShortLink(ctx, shortLink)
	if err != nil {
		return Link{}, apperr.NewSlugError(err.Error(), "not-found")
	}

	_ = s.reportSvc.LinkClicked(ctx, shortLink, time.Now())

	return link, nil
}
