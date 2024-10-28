package repository

import (
	"context"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/link/app/service"
	"github.com/jmoiron/sqlx"
)

type LinkRepository struct {
	db *sqlx.DB
}

func NewLinkRepository(db *sqlx.DB) LinkRepository {
	return LinkRepository{db: db}
}

func (r LinkRepository) Create(ctx context.Context, originalLink, shortLink string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO links (original_link, short_link) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, originalLink, shortLink)
	if err != nil {
		return err
	}

	return nil
}

func (r LinkRepository) GetByShortLink(ctx context.Context, shortLink string) (service.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM links WHERE short_link=$1`

	link := service.Link{}
	err := r.db.QueryRowxContext(ctx, query, shortLink).
		Scan(
			&link.ID,
			&link.OriginalLink,
			&link.ShortLink,
			&link.CreatedAt,
		)
	if err != nil {
		return service.Link{}, err
	}

	return link, nil
}
