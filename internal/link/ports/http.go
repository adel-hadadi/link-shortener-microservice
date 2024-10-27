package ports

import (
	"log"
	"net/http"

	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app: app}
}

const (
	OriginalLinkRequired = "original link is required"
)

func (h HttpServer) CreateLink(w http.ResponseWriter, r *http.Request) {
	req := CreateLink{}

	err := render.Decode(r, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if req.OriginalLink == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(OriginalLinkRequired))
		return
	}

	log.Printf("create link => %s\n", req.OriginalLink)
}
