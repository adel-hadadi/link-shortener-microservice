package main

import _ "github.com/adel-hadadi/link-shotener/internal/common/setup"

import (
	"fmt"

	"github.com/adel-hadadi/link-shotener/internal/common/genproto/report"
	"github.com/adel-hadadi/link-shotener/internal/common/server"
	"github.com/adel-hadadi/link-shotener/internal/report/ports"
	"github.com/adel-hadadi/link-shotener/internal/report/service"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	ErrLoadEnv = "error on loading .env file: %w"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf(ErrLoadEnv, err))
	}

	app, cleanup := service.NewApplication()
	defer cleanup()

	ports.RunScheduler(app)

	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(app)
		report.RegisterReportServiceServer(server, svc)
	})

}
