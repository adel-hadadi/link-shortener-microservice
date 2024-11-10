package service

import (
	"github.com/adel-hadadi/link-shotener/internal/link/adapters"
	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/adel-hadadi/link-shotener/internal/link/app/command"
	"github.com/adel-hadadi/link-shotener/internal/link/app/query"
	"github.com/jmoiron/sqlx"
)

func NewApplication(db *sqlx.DB) app.Application {
	linkRepo := adapters.NewPostgresLinkRepository(db)

	return app.Application{
		Commands: app.Commands{
			GenerateShortURL: command.NewGenerateShortURLHandler(linkRepo),
		},
		Queries: app.Queries{
			RetrieveOriginalURL: query.NewRetrieveOriginalURLHandler(linkRepo),
			RetrieveShortURL:    query.NewRetrieveShortURLHandler(linkRepo),
		},
	}
}
