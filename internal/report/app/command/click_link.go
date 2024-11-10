package command

import (
	"context"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/report/domain/report"
)

type LinkClickHandler struct {
	repo report.Repository
}

type LinkClick struct {
	Link      string
	ClickedAt time.Time
}

func NewLinkClickHandler(repo report.Repository) LinkClickHandler {
	return LinkClickHandler{repo: repo}
}

func (h LinkClickHandler) Handle(ctx context.Context, cmd LinkClick) error {
	return h.repo.Create(ctx, cmd.Link, cmd.ClickedAt)
}
