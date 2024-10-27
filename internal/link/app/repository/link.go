package repository

import (
	"context"
	"time"

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
