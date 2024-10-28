package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type ClickRepository struct {
	db *sqlx.DB
}

func NewClickRepository(db *sqlx.DB) ClickRepository {
	return ClickRepository{db}
}

func (r ClickRepository) Create(ctx context.Context, linkID int, ipAddress string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO clicks (link_id, ip_address) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, linkID, ipAddress)

	return err
}
