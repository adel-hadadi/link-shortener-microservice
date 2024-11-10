package adapters

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresLinkRepository struct {
	db *sqlx.DB
}

func NewPostgresLinkRepository(db *sqlx.DB) PostgresLinkRepository {
	return PostgresLinkRepository{db: db}
}

func (r PostgresLinkRepository) Create(ctx context.Context, originalLink, shortLink string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO links (original_link, short_link) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, originalLink, shortLink)
	return err
}

func (r PostgresLinkRepository) GetByOriginalURL(ctx context.Context, originalLink string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT short_link FROM links WHERE original_link=$1`

	var shortLink string
	err := r.db.QueryRowxContext(ctx, query, originalLink).
		Scan(
			&shortLink,
		)

	return shortLink, err
}

func (r PostgresLinkRepository) GetByShortURL(ctx context.Context, shortLink string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT original_link FROM links WHERE short_link=$1`

	var originalLink string
	err := r.db.QueryRowxContext(ctx, query, shortLink).
		Scan(
			&originalLink,
		)

	return originalLink, err
}
