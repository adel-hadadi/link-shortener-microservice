package app

import "github.com/adel-hadadi/link-shotener/internal/link/app/service"

type Application struct {
	Services Services
}

type Services struct {
	LinkService  service.LinkService
	ClickService service.ClickService
}
