package service

import (
	"fmt"

	grpcClient "github.com/adel-hadadi/link-shotener/internal/common/client"
	"github.com/adel-hadadi/link-shotener/internal/link/adapters"
	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/adel-hadadi/link-shotener/internal/link/app/repository"
	"github.com/adel-hadadi/link-shotener/internal/link/app/service"
	"github.com/jmoiron/sqlx"
)

func NewApplication(db *sqlx.DB) (app.Application, func()) {
	reportClient, closeReportClient, err := grpcClient.NewReportClient()
	if err != nil {
		panic(fmt.Errorf("error on connecting to report gRPC: %w", err))
	}

	reportGrpc := adapters.NewReportGrpc(reportClient)

	linkRepo := repository.NewLinkRepository(db)

	return app.Application{
			Services: app.Services{
				LinkService: service.NewLinkService(linkRepo, reportGrpc),
			},
		}, func() {
			_ = closeReportClient()
		}
}
