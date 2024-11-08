package main

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
	defer func() {
		_ = closeLinkClient()
	}()

	linkGrpc := adapters.NewLinkGrpc(client)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return HandlerFromMux(NewHttpServer(linkGrpc), router)
	})
}
