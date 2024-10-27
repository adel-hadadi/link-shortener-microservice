package main

import (
	"net/http"

	"github.com/adel-hadadi/link-shotener/internal/common/server"
	"github.com/adel-hadadi/link-shotener/internal/link/ports"
	"github.com/adel-hadadi/link-shotener/internal/link/service"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := service.NewApplication()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(app),
			router,
		)
	})
}
