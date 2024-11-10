package report

import (
	"context"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/report/app/query"
)

type Repository interface {
	Create(ctx context.Context, shortURL string, clickedAt time.Time) error
	GetLastHourClicks(ctx context.Context) ([]query.ClickCount, error)
}
