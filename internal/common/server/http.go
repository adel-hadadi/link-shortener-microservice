package server

import (
	"log"
	"net/http"
	"os"

	"github.com/adel-hadadi/link-shotener/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddleware(apiRouter)

	rootRouter := chi.NewRouter()

	rootRouter.Mount("/api", createHandler(apiRouter))

	log.Println("Starting HTTP server")

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		panic(err)
	}
}

func setMiddleware(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Logger)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
}
