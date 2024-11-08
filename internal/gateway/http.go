package main

import (
	"net/http"

	"github.com/adel-hadadi/link-shotener/internal/common/server/httperr"
	"github.com/adel-hadadi/link-shotener/internal/common/server/httpres"
	"github.com/go-chi/render"
)

type HttpServer struct {
	linkService LinkService
}

func NewHttpServer(linkService LinkService) HttpServer {
	return HttpServer{linkService: linkService}
}

func (h HttpServer) CreateLink(w http.ResponseWriter, r *http.Request) {
	req, err := httpres.Bind[CreateLink](r)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	shortLink, err := h.linkService.CreateLink(r.Context(), req.OriginalLink)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)

	render.JSON(w, r, Link{
		OriginalLink: req.OriginalLink,
		ShortLink:    shortLink,
	})
}

func (h HttpServer) RedirectToURL(w http.ResponseWriter, r *http.Request, shortURL string) {
	link, err := h.linkService.GetLink(r.Context(), shortURL)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	http.Redirect(w, r, link, 302)
}
