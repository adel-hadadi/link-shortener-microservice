package service

import (
	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/adel-hadadi/link-shotener/internal/link/app/repository"
	"github.com/adel-hadadi/link-shotener/internal/link/app/service"
	"github.com/jmoiron/sqlx"
)

func NewApplication(db *sqlx.DB) app.Application {
	linkRepository := repository.NewLinkRepository(db)

	return app.Application{
		Services: app.Services{
			LinkService: service.NewLinkService(linkRepository),
		},
	}
}
