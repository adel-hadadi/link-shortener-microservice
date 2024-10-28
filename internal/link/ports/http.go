package ports

import (
	"net/http"

	"github.com/adel-hadadi/link-shotener/internal/common/server/httperr"
	"github.com/adel-hadadi/link-shotener/internal/link/app"
	"github.com/adel-hadadi/link-shotener/internal/link/app/service"
	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app: app}
}

func (h HttpServer) CreateLink(w http.ResponseWriter, r *http.Request) {
	req := CreateLink{}

	err := render.Decode(r, &req)
	if err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	if req.OriginalLink == "" {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	err = h.app.Services.LinkService.Create(r.Context(), service.CreateLink{
		OriginalLink: req.OriginalLink,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h HttpServer) RedirectLink(w http.ResponseWriter, r *http.Request, shortLink string) {
	link, err := h.app.Services.LinkService.Get(r.Context(), shortLink)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.app.Services.ClickService.Create(r.Context(), service.CreateClick{LinkID: link.ID, IPAddress: r.RemoteAddr})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	http.Redirect(w, r, link.OriginalLink, http.StatusFound)
}
