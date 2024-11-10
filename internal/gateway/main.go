package main

import _ "github.com/adel-hadadi/link-shotener/internal/common/setup"

import (
	"fmt"
	"net/http"

	grpcClient "github.com/adel-hadadi/link-shotener/internal/common/client"
	"github.com/adel-hadadi/link-shotener/internal/common/server"
	"github.com/adel-hadadi/link-shotener/internal/gateway/adapters"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const (
	ErrLoadEnv = "error on loading .env file: %w"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf(ErrLoadEnv, err))
	}

	client, closeLinkClient, err := grpcClient.NewLinkClient()
	if err != nil {
		panic(err)
	}
	defer closeLinkClient()

	linkGrpc := adapters.NewLinkGrpc(client)

	reportClient, closeReportClient, err := grpcClient.NewReportClient()
	if err != nil {
		panic(err)
	}
	defer closeReportClient()

	reportGrpc := adapters.NewReportGrpc(reportClient)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return HandlerFromMux(NewHttpServer(linkGrpc, reportGrpc), router)
	})
}
