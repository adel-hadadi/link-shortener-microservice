package app

import (
	"github.com/adel-hadadi/link-shotener/internal/report/app/command"
	"github.com/adel-hadadi/link-shotener/internal/report/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	GenerateReport command.GenerateReportHandler
	LinkClick      command.LinkClickHandler
}

type Queries struct {
	DownloadReport query.DownloadReportHandler
}
