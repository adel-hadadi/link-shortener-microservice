package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/adel-hadadi/link-shotener/internal/common/server/httperr"
	"github.com/adel-hadadi/link-shotener/internal/common/server/httpres"
	"github.com/go-chi/render"
)

type HttpServer struct {
	linkService   LinkService
	reportService ReportService
}

func NewHttpServer(linkService LinkService, reportService ReportService) HttpServer {
	return HttpServer{linkService: linkService, reportService: reportService}
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

	go func() {
		err := h.reportService.LinkClicked(r.Context(), shortURL, time.Now())
		if err != nil {
			log.Println(fmt.Errorf("error on calling link clicked event: %w", err))
		}
	}()

	http.Redirect(w, r, link, 302)
}

func (h HttpServer) DownloadExcelFile(w http.ResponseWriter, r *http.Request, fileName string) {
	fileName = fmt.Sprintf("clicks-report-%s.xlsx", fileName)

	stream, err := h.reportService.DownloadReport(r.Context(), fileName)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	// Set headers for Excel file download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			httperr.RespondWithSlugError(fmt.Errorf("Failed to receive chunk: %w", err), w, r)
			return
		}

		if _, err := w.Write(chunk.Content); err != nil {
			httperr.RespondWithSlugError(fmt.Errorf("Failed to receive chunk: %w", err), w, r)
			return
		}
	}
}
