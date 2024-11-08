package app

import "github.com/adel-hadadi/link-shotener/internal/report/app/service"

type Application struct {
	Services Services
}

type Services struct {
	ReportService service.ReportService
}
